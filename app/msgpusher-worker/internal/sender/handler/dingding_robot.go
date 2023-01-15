package handler

import (
	"austin-v2/app/msgpusher-common/domain/account"
	"austin-v2/app/msgpusher-common/domain/content_model"
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/data"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/accountHelper"
	"austin-v2/pkg/utils/arrayHelper"
	"austin-v2/pkg/utils/contentHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/ants/v2"
	"github.com/wanghuiyt/ding"
	"strings"
)

type DingDingRobotHandler struct {
	BaseHandler
	logger *log.Helper
	rds    redis.Cmdable
	sc     *biz.SendAccountUseCase
	mrr    data.IMsgRecordRepo
}

const SendAll = "@all"

func NewDingDingRobotHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
	mrr data.IMsgRecordRepo,
) *DingDingRobotHandler {
	return &DingDingRobotHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/ding-ding-robot")),
		rds:    rds,
		sc:     sc,
		mrr:    mrr,
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
	record := h.getRecord(taskInfo, strings.Join(taskInfo.Receiver, ","))
	record.Channel = h.Name()
	if err = d.SendMessage(content.Content, at...); err != nil {
		h.logger.WithContext(ctx).Errorw(
			"msg", "dingDingRobotHandler send err",
			"err", err,
			"requestId", taskInfo.RequestId)
		record.Msg = "推送失败: " + err.Error()
	} else {
		h.logger.WithContext(ctx).Infow(
			"msg", "dingDingRobotHandler send success",
			"requestId", taskInfo.RequestId)
		record.Msg = "推送成功"
	}
	_ = ants.Submit(func() {
		_ = h.mrr.InsertMany(ctx, []interface{}{record})
	})
	return nil
}

func (h *DingDingRobotHandler) Name() string {
	return channelType.TypeCodeEn[channelType.DingDingRobot]
}
