package handler

import "context"

type BaseHandler struct {
}

func (b BaseHandler) Allow(ctx context.Context) bool {
	return true
}
