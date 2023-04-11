package middlewarex

import (
	mgrApi "austin-v2/api/mgr"
	"austin-v2/app/mgr/internal/data"
	"austin-v2/app/mgr/internal/pkg/ctxdata"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"strings"
)

// 不需要登录的部分
var noNeedLogin = map[string]struct{}{
	mgrApi.OperationSystemAdminLogin:     {},
	mgrApi.OperationCommonGetIndexConfig: {},
}

var ErrUnauthorized = errors.New(401, "UNAUTHORIZED_INFO_MISSING", "授权已过期或授权异常,请重新授权")

func NewLoginWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := noNeedLogin[operation]; ok {
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
					return nil, ErrUnauthorized
				}
				jwtToken := auths[1]
				token, _ := data.RedisCli.Get(ctx, jwtToken).Result()
				if token == "" {
					return nil, ErrUnauthorized
				}
				// 设置 Authorization
				tr.RequestHeader().Set("Authorization", "Bearer "+token)
				tr.RequestHeader().Set("X-Token", jwtToken)
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
				return nil, ErrUnauthorized
			}
			claimInfo := claim.(jwt2.MapClaims)
			ctx = ctxdata.SetMgrID(ctx, cast.ToInt32(claimInfo["uid"]))
			ctx = ctxdata.SetMgrName(ctx, cast.ToString(claimInfo["uname"]))
			return handler(ctx, req)
		}
	}
}

func Jwt(ApiKey string) middleware.Middleware {
	// 对于需要登录的路由进行jwt中间件验证
	return selector.Server(
		// 设置header信息
		setHeaderInfo(),
		// 解析jwt
		jwt.Server(func(token *jwt2.Token) (interface{}, error) {
			return []byte(ApiKey), nil
		},
			jwt.WithSigningMethod(jwt2.SigningMethodHS256),
			jwt.WithClaims(func() jwt2.Claims {
				return jwt2.MapClaims{}
			})),
		// 设置全局ctx
		setUserInfo(),
	).
		Match(NewLoginWhiteListMatcher()).
		Build()
}
