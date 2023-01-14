package handler

import (
	"austin-v2/pkg/types"
	"context"
)

type BaseHandler struct {
}

// Allow 限流方法 默认不限流
func (b BaseHandler) Allow(_ context.Context, _ *types.TaskInfo) bool {
	return true
}
