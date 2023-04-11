package metaHelper

import (
	"austin-v2/utils/jsonHelper"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/metadata"
)

type LoginUser struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

const loginKey = "x-md-global-admin-login"

func WithMetaAdminUser(ctx context.Context, user LoginUser) context.Context {
	return metadata.AppendToClientContext(ctx, loginKey, jsonHelper.MustToString(user))
}

func GetMetaAdminUser(ctx context.Context) LoginUser {
	var res LoginUser
	if serverContext, ok := metadata.FromServerContext(ctx); ok {
		_ = json.Unmarshal([]byte(serverContext.Get(loginKey)), &res)
	}
	return res
}
