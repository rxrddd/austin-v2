package handler

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/common/domain/content_model"
	"austin-v2/common/enums/channelType"
	"austin-v2/pkg/types"
	"austin-v2/utils/contentHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type DingDingWorkNoticeHandler struct {
	BaseHandler
	logger *log.Helper
	rds    redis.Cmdable
	sc     *biz.SendAccountUseCase
}

func NewDingDingWorkNoticeHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
) *DingDingWorkNoticeHandler {
	return &DingDingWorkNoticeHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/ding-ding-notice")),
		rds:    rds,
		sc:     sc,
	}
}

func (h DingDingWorkNoticeHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	contentHelper.GetContentModel(taskInfo.ContentModel, &content)
	h.logger.WithContext(ctx).Infow(
		"msg", "DingDingWorkNoticeHandler send success",
		"requestId", taskInfo.RequestId)

	return nil
}

func (h *DingDingWorkNoticeHandler) Name() string {
	return channelType.TypeCodeEn[channelType.DingDingWorkNotice]
}
