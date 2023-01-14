package handler

import (
	"austin-v2/app/msgpusher-common/domain/account"
	"austin-v2/app/msgpusher-common/domain/content_model"
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/accountHelper"
	"austin-v2/pkg/utils/arrayHelper"
	"austin-v2/pkg/utils/contentHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/wanghuiyt/ding"
)

type DingDingRobotHandler struct {
	BaseHandler
	logger *log.Helper
	rds    redis.Cmdable
	sc     *biz.SendAccountUseCase
}

const SendAll = "@all"

func NewDingDingRobotHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
) *DingDingRobotHandler {
	return &DingDingRobotHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/ding-ding-robot")),
		rds:    rds,
		sc:     sc,
	}
}
func (h *DingDingRobotHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	var content content_model.DingDingContentModel
	contentHelper.GetContentModel(taskInfo.ContentModel, &content)

	var acc account.DingDingRobotAccount
	if err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc); err != nil {
		h.logger.WithContext(ctx).Errorw(
			"msg", "dingDingRobotHandler get account err",
			"err", err,
			"requestId", taskInfo.RequestId)
		return err
	}
	var at []string
	d := ding.Webhook{
		AccessToken: acc.AccessToken,
		Secret:      acc.Secret,
		EnableAt:    true,
	}

	if arrayHelper.ArrayStringIn(taskInfo.Receiver, SendAll) {
		d.AtAll = true
	} else {
		at = taskInfo.Receiver
	}

	if err = d.SendMessage(content.Content, at...); err != nil {
		h.logger.WithContext(ctx).Errorw(
			"msg", "dingDingRobotHandler send err",
			"err", err,
			"requestId", taskInfo.RequestId)
		return err
	}
	h.logger.WithContext(ctx).Infow(
		"msg", "dingDingRobotHandler send success",
		"requestId", taskInfo.RequestId)

	return nil
}

func (h *DingDingRobotHandler) Name() string {
	return channelType.TypeCodeEn[channelType.DingDingRobot]
}
