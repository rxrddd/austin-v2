package server

import (
	"context"
	"encoding/json"
	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/conf"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/data"
	"github.com/ZQCard/kratos-base-project/app/project/admin/internal/service"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/ZQCard/kratos-base-project/pkg/middleware/casbin"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	stdHttp "net/http"
	"strings"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.admin.v1.Admin/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// 设置header信息
func setHeaderInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			if tr, ok := transport.FromServerContext(ctx); ok {
				// 将请求信息放入ctx中
				if ht, ok := tr.(*http.Transport); ok {
					ctx = context.WithValue(ctx, "RemoteAddr", ht.Request().RemoteAddr)
				}

				// 将md5格式的Authorization置换redis中真正的jwt值
				auths := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
					return nil, errResponse.SetErrByReason(errResponse.ReasonAdministratorUnauthorized)
				}
				jwtToken := auths[1]
				token, _ := data.RedisCli.Get(jwtToken).Result()
				if token == "" {
					return nil, errResponse.SetErrByReason(errResponse.ReasonAdministratorUnauthorized)
				}
				// 设置 Authorization
				tr.RequestHeader().Set("Authorization", "Bearer "+token)
			}
			return handler(ctx, req)
		}
	}
}

// 设置用户信息
func setUserInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedInfoMissing)
			}
			claimInfo := claim.(jwt2.MapClaims)
			AdministratorId := int64(claimInfo["AdministratorId"].(float64))
			ctx = context.WithValue(ctx, "kratos-AdministratorId", AdministratorId)
			ctx = context.WithValue(ctx, "kratos-AdministratorUsername", claimInfo["AdministratorUsername"])
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new a HTTP serviceName.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, service *service.AdminInterface, tp *tracesdk.TracerProvider, logger log.Logger) *http.Server {
	// 初始化基础数据库 casbin权限控制策略,连接基础库
	db, err := gorm.Open(mysql.Open(ac.CasbinSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql connect error: %v", err)
	}
	a, _ := gormadapter.NewAdapterByDB(db)
	// 加载权限配置文件
	m, err := model.NewModelFromString(ac.CasbinModel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Client(),
			// 日志记录
			logging.Server(logger),

			// 对于需要登录的路由进行jwt中间件验证
			selector.Server(
				// 设置header信息
				setHeaderInfo(),
				// 解析jwt
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				},
					jwt.WithSigningMethod(jwt2.SigningMethodHS256),
					jwt.WithClaims(func() jwt2.Claims {
						return jwt2.MapClaims{}
					})),
				// 设置全局ctx
				setUserInfo(),
				// 权限中间件
				casbin.Server(
					casbin.WithCasbinModel(m),
					casbin.WithCasbinPolicy(a),
				),
			).
				Match(NewWhiteListMatcher()).
				Build(),
		),
		// 跨域设置
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "AccessToken", "X-Token", "Accept"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// 增加成功自定义json返回值
	opts = append(opts, http.ResponseEncoder(responseEncoder))
	// 增加错误自定义json返回值
	opts = append(opts, http.ErrorEncoder(errorEncoder))

	srv := http.NewServer(opts...)
	route := srv.Route("/")
	route.POST("/files/v1/uploadFile", service.UploadFile)
	v1.RegisterAdminHTTPServer(srv, service)
	return srv
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func responseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := &Response{}
	reply.Code = 0
	reply.Message = "success"

	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(v)
	_ = json.Unmarshal(data, &reply.Data)
	if err != nil {
		return err
	}

	data, err = codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", codec.Name())
	w.WriteHeader(stdHttp.StatusOK)
	w.Write(data)
	return nil
}

func errorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// 返回码均是200
	w.WriteHeader(stdHttp.StatusOK)
	// 重写errResponse
	err = errResponse.SetCustomizeErrInfo(err)
	se := errors.FromError(err)
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, _ = w.Write(body)
}
