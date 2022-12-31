package handler

import (
	"austin-v2/app/msgpusher-worker/internal/enums/channelType"
	"austin-v2/app/msgpusher-worker/internal/svc"
	"austin-v2/pkg/types"
	"context"
	"fmt"
)

type smsHandler struct {
	svcCtx *svc.ServiceContext
	BaseHandler
}

func NewSmsHandler(svcCtx *svc.ServiceContext) types.IHandler {
	return &smsHandler{
		svcCtx: svcCtx,
	}
}

func (h *smsHandler) Name() string {
	return channelType.TypeCodeEn[channelType.Sms]
}

func (h *smsHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	fmt.Println("sms sender")
	return nil
}
