package handler

import (
	"austin-v2/pkg/types"
	"context"
)

type BaseHandler struct {
}

func (b BaseHandler) Allow(ctx context.Context, taskInfo *types.TaskInfo) bool {
	return true
}
