package server

import (
	"context"
	"fmt"
	"github.com/ZQCard/kratos-base-project/third_party/errResponse"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	stdhttp "net/http"
	"time"

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

	"github.com/ZQCard/kratos-base-project/api/project/admin/v1"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/conf"
	"github.com/ZQCard/kratos-base-project/app/project/admin/service/internal/service"
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

// 获取当前服务operation
func getOperation() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				fmt.Println(tr.Operation())
			}
			return handler(ctx, req)
		}
	}
}

// 设置全局参数
func setGlobalData() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return handler(ctx, req)
			}
			claimInfo := claim.(jwt2.MapClaims)
			AdministratorId := int64(claimInfo["AdministratorId"].(float64))
			ctx = context.WithValue(ctx, "kratos-AdministratorId", AdministratorId)
			ctx = context.WithValue(ctx, "kratos-AdministratorUsername", claimInfo["AdministratorUsername"])
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, service *service.AdminInterface, tp *tracesdk.TracerProvider, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Client(),
			// 日志记录
			logging.Server(logger),
			// 对于需要登录的路由进行jwt中间件验证
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				},
					jwt.WithSigningMethod(jwt2.SigningMethodHS256),
					jwt.WithClaims(func() jwt2.Claims {
						return jwt2.MapClaims{}
					})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
			setGlobalData(),
		),
		// 跨域设置
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "AccessToken", "X-Token", "Accept"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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
	v1.RegisterAdminHTTPServer(srv, service)
	return srv
}

func responseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, i interface{}) error {
	type response struct {
		Code    int64       `json:"code"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Ts      int64       `json:"ts"`
	}
	reply := &response{
		Code: 0,
		Data: i,
		Ts:   time.Now().Unix(),
	}
	codec := encoding.GetCodec("json")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return nil
}

func errorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {

	codec, _ := http.CodecForRequest(r, "Accept")
	w.Header().Set("Content-Type", "application/"+codec.Name())
	// 返回码均是200
	w.WriteHeader(stdhttp.StatusOK)
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
