package handler

import (
	"austin-v2/app/msgpusher-common/domain/account"
	"austin-v2/app/msgpusher-common/domain/content_model"
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/data"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/accountHelper"
	"austin-v2/pkg/utils/contentHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/panjf2000/ants/v2"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/spf13/cast"
	"strings"
)

const colorSep = "|" //以|分割颜色

//公众号订阅消息
type OfficialAccountHandler struct {
	BaseHandler
	logger *log.Helper
	rds    redis.Cmdable
	sc     *biz.SendAccountUseCase
	mrr    data.IMsgRecordRepo
}

func NewOfficialAccountHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
	mrr data.IMsgRecordRepo,
) *OfficialAccountHandler {
	return &OfficialAccountHandler{
		logger: log.NewHelper(log.With(logger, "module", "sender/sms")),
		rds:    rds,
		sc:     sc,
		mrr:    mrr,
	}
}
func (h *OfficialAccountHandler) Name() string {
	return channelType.TypeCodeEn[channelType.OfficialAccounts]
}

func (h *OfficialAccountHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	var content content_model.OfficialAccountsContentModel
	contentHelper.GetContentModel(taskInfo.ContentModel, &content)
	//拼接消息发送
	var acc account.OfficialAccount
	if err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc); err != nil {
		return err
	}
	wc := wechat.NewWechat()
	cacheImpl := cache.NewMemory()

	cfg := &offConfig.Config{
		AppID:          acc.AppID,
		AppSecret:      acc.AppSecret,
		Token:          acc.Token,
		EncodingAESKey: acc.EncodingAESKey,
		Cache:          cacheImpl,
	}
	subscribe := wc.GetOfficialAccount(cfg).GetTemplate()
	templateSn := content.TemplateSn
	url := content.Url
	params := make(map[string]*message.TemplateDataItem, len(content.Data))
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
		params[key] = &message.TemplateDataItem{Value: value, Color: color}
	}
	var (
		msgIds  []int64
		records []interface{}
	)
	for _, receiver := range taskInfo.Receiver {
		msgID, err := subscribe.Send(&message.TemplateMessage{
			ToUser:     receiver,
			TemplateID: templateSn,
			URL:        url,
			Data:       params,
			MiniProgram: struct {
				AppID    string `json:"appid"`
				PagePath string `json:"pagepath"`
			}(struct {
				AppID    string
				PagePath string
			}{AppID: content.MiniProgram.Appid, PagePath: content.MiniProgram.PagePath}),
		})
		record := h.getRecord(taskInfo, receiver)
		record.MsgId = cast.ToString(msgID)
		record.Channel = h.Name()

		if err != nil {
			h.logger.WithContext(ctx).Errorw(
				"msg", "OfficialAccountHandler send msg",
				"err", err,
				"receiver", receiver,
				"templateSn", templateSn)
			record.Msg = "推送失败: " + err.Error()
		} else {
			msgIds = append(msgIds, msgID)
			record.Msg = "推送成功"
		}
		records = append(records, record)
	}
	if len(msgIds) > 0 {
		h.logger.WithContext(ctx).Infow(
			"msg", "OfficialAccountHandler send success",
			"requestId", taskInfo.RequestId,
			"msgIds", msgIds)
	}
	_ = ants.Submit(func() {
		_ = h.mrr.InsertMany(ctx, records)
	})
	return nil
}
