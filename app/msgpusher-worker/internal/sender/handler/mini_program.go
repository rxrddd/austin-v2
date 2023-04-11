package handler

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/data"
	"austin-v2/common/domain/account"
	"austin-v2/common/domain/content_model"
	"austin-v2/common/enums/channelType"
	"austin-v2/common/model"
	"austin-v2/pkg/types"
	"austin-v2/utils/accountHelper"
	"austin-v2/utils/contentHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniprogramConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"strings"
)

type MiniProgramHandler struct {
	BaseHandler
	logger *log.Helper
	rds    redis.Cmdable
	sc     *biz.SendAccountUseCase
	mrr    data.IMsgRecordRepo
}

func NewMiniProgramHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
	mrr data.IMsgRecordRepo,
) *MiniProgramHandler {
	return &MiniProgramHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/mini-program")),
		rds:    rds,
		sc:     sc,
		mrr:    mrr,
	}
}

func (h *MiniProgramHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	var content content_model.MiniProgramContentModel
	contentHelper.GetContentModel(taskInfo.ContentModel, &content)
	//拼接消息发送
	var acc account.OfficialAccount

	if err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc); err != nil {
		return errors.Wrap(err, "OfficialAccountHandler get account err")
	}
	wc := wechat.NewWechat()
	cacheImpl := cache.NewMemory()

	cfg := &miniprogramConfig.Config{
		AppID:     acc.AppID,
		AppSecret: acc.AppSecret,
		Cache:     cacheImpl,
	}
	sub := wc.GetMiniProgram(cfg).GetSubscribe()
	templateSn := content.TemplateSn
	params := make(map[string]*subscribe.DataItem, len(content.Data))
	for key, val := range content.Data {
		color := ""
		value := ""
		arr := strings.Split(val, colorSep)
		if len(arr) == 1 {
			value = arr[0]
		}
		if len(arr) == 2 {
			value = arr[0]
			color = arr[1]
		}
		params[key] = &subscribe.DataItem{Value: value, Color: color}
	}
	var records []*model.MsgRecord
	for _, receiver := range taskInfo.Receiver {
		err := sub.Send(&subscribe.Message{
			ToUser:           receiver,
			TemplateID:       templateSn,
			Page:             content.Page,
			Data:             params,
			MiniprogramState: content.MiniProgramState,
			Lang:             content.Lang,
		})
		record := h.getRecord(taskInfo, receiver)
		record.Channel = h.Name()

		if err != nil {
			h.logger.WithContext(ctx).Errorw(
				"msg", "MiniProgramHandler send err",
				"err", err,
				"receiver", receiver,
				"templateSn", templateSn)
			record.Msg = "推送失败: " + err.Error()
		} else {
			record.Msg = "推送成功"
		}
		records = append(records, record)
	}
	h.logger.WithContext(ctx).Infow(
		"msg", "MiniProgramHandler send success",
		"requestId", taskInfo.RequestId)
	_ = ants.Submit(func() {
		_ = h.mrr.InsertMany(ctx, records)
	})
	return nil
}

func (h *MiniProgramHandler) Name() string {
	return channelType.TypeCodeEn[channelType.DingDingWorkNotice]
}
