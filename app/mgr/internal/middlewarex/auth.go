package middlewarex

import (
	mgrApi "austin-v2/api/mgr"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strings"
)

var ErrUnAuth = errors.New(401, "UNAUTH", "你没有该功能的权限")

// 不需要鉴权的路径
var noNeedAuth = map[string]struct{}{
	mgrApi.OperationSystemAdminLogin:  {},
	mgrApi.OperationSystemAdminLogout: {},
	mgrApi.OperationSystemGetSelfInfo: {},
}

func NewAuthWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := noNeedAuth[operation]; ok {
			return false
		}
		return true
	}
}
func authMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(*http.Transport); ok {
					authStr := strings.TrimLeft(strings.ReplaceAll(ht.Request().RequestURI, "/", ":"), ":")
					fmt.Println(authStr)
				}
			}
			return handler(ctx, req)
		}
	}
}

func Auth() middleware.Middleware {
	return selector.Server(
		// Auth验证
		authMiddleware(),
	).
		Match(NewAuthWhiteListMatcher()).
		Build()
}
