package metaHelper

import (
	"austin-v2/pkg/utils/jsonHelper"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/metadata"
)

type LoginUser struct {
	UserId   int64
	UserName string
}

const LoginKey = "x-md-global-admin-login"

func WithContext(ctx context.Context, user LoginUser) context.Context {
	return WithClientContext(ctx, LoginKey, jsonHelper.MustToString(user))
}
func WithClientContext(ctx context.Context, kv ...string) context.Context {
	ctx = metadata.AppendToClientContext(ctx, kv...)
	return ctx
}

func GetMetaAdminUser(ctx context.Context) LoginUser {
	serverContext, _ := metadata.FromServerContext(ctx)
	var res LoginUser
	_ = json.Unmarshal([]byte(serverContext.Get(LoginKey)), &res)
	return res

}
