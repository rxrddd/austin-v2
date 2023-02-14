package metaHelper

import (
	"austin-v2/app/project/admin/pkg/ctxdata"
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
)

func MetaUserMiddleware() func(handler middleware.Handler) middleware.Handler {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = WithMetaAdminUser(ctx, LoginUser{
				UserId:   ctxdata.GetAdminId(ctx),
				UserName: ctxdata.GetAdminName(ctx),
			})
			return handler(ctx, req)
		}
	}
}
