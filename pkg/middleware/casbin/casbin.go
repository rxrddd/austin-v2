package casbin

import (
	"context"
	"fmt"
	"github.com/ZQCard/kratos-base-project/pkg/errResponse"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

type Option func(*options)

const CabinObj = "AdministratorRole"

type options struct {
	model    model.Model
	policy   persist.Adapter
	enforcer *casbin.SyncedEnforcer
}

func WithCasbinModel(model model.Model) Option {
	return func(o *options) {
		o.model = model
	}
}

func WithCasbinPolicy(policy persist.Adapter) Option {
	return func(o *options) {
		o.policy = policy
	}
}

func Server(opts ...Option) middleware.Middleware {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	o.enforcer, _ = casbin.NewSyncedEnforcer(o.model, o.policy)

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if o.enforcer == nil {
				return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedInfoMissing)
			}
			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedInfoMissing)
			}
			claimInfo := claim.(jwt2.MapClaims)
			if claimInfo[CabinObj] == nil {
				return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedInfoMissing)
			}
			role := claimInfo[CabinObj].(string)

			// 获取当前服务operation，验证策略为 operation + method + role(超级管理员不做判断)
			if tr, ok := transport.FromServerContext(ctx); ok && role != "超级管理员" {
				// 获取请求方法
				act := tr.RequestHeader().Get("Http-Method")
				// 获取请求的PATH
				obj := tr.Operation()
				fmt.Println("tr")
				fmt.Println(tr)
				// 权限判断
				allowed, err := o.enforcer.Enforce(role, obj, act)
				if err != nil {
					return nil, err
				}
				if !allowed {
					return nil, errResponse.SetErrByReason(errResponse.ReasonUnauthorizedRole)
				}
			}
			return handler(ctx, req)
		}
	}
}
