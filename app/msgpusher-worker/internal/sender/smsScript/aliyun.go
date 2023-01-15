package smsScript

import (
	"austin-v2/app/msgpusher-common/domain/account"
	"austin-v2/app/msgpusher-common/domain/content_model"
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/pkg/mq"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/accountHelper"
	"austin-v2/pkg/utils/contentHelper"
	"austin-v2/pkg/utils/jsonHelper"
	"austin-v2/pkg/utils/stringHelper"
	"austin-v2/pkg/utils/timeHelper"
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	smsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cast"
	"time"
)

type AliyunSms struct {
	logger   *log.Helper
	mqHelper mq.IMessagingClient
	sc       *biz.SendAccountUseCase
}

func NewAliyunSms(
	logger log.Logger,
	mqHelper mq.IMessagingClient,
	sc *biz.SendAccountUseCase,
) *AliyunSms {
	return &AliyunSms{
		logger:   log.NewHelper(log.With(logger, "module", "sender/smsScript/yunpian")),
		mqHelper: mqHelper,
		sc:       sc,
	}
}
func (h *AliyunSms) Name() string {
	return "aliyun"
}
func (h *AliyunSms) Send(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	var acc account.AliyunSmsAccount
	if err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc); err != nil {
		return err
	}
	config := &openapi.Config{
		AccessKeyId:     &acc.AccessKeyId,
		AccessKeySecret: &acc.AccessSecret,
	}
	// 访问的域名
	config.Endpoint = tea.String(acc.GatewayURL)
	cli, err := smsapi.NewClient(config)
	if err != nil {
		return fmt.Errorf("smsapi.NewClient error = %v", err)
	}
	var content content_model.SmsContentModel
	contentHelper.GetContentModel(taskInfo.ContentModel, &content)
	records := make([]model.SmsRecord, 0)
	for _, receiver := range taskInfo.Receiver {
		request := &smsapi.SendSmsRequest{}
		request.SetPhoneNumbers(receiver)
		request.SetSignName(acc.SignName)
		request.SetTemplateCode(taskInfo.TemplateSn)
		request.SetTemplateParam(jsonHelper.MustToString(taskInfo.MessageParam.Variables))

		var response *smsapi.SendSmsResponse
		if response, err = cli.SendSms(request); err != nil {
			h.logger.WithContext(ctx).Errorw(
				"msg", "Client.Send() error",
				"err", err,
				"receiver", receiver,
				"request", request.String())
			continue
		}
		if *response.Body.Code == "OK" {
			records = append(records, h.smsRecord(response, taskInfo, receiver, content))
		}
	}
	_ = ants.Submit(func() {
		if err = h.mqHelper.Publish(jsonHelper.MustToByte(records), "sms.record"); err != nil {
			h.logger.WithContext(ctx).Errorw("msg", "aliyun send publish err", "err", err)
		}
	})
	return nil
}
func (h *AliyunSms) smsRecord(response *smsapi.SendSmsResponse,
	info *types.TaskInfo,
	phoneNumber string,
	content content_model.SmsContentModel,
) model.SmsRecord {
	var insert = model.SmsRecord{
		ID:                stringHelper.NextID(),
		RequestID:         info.RequestId,
		MessageTemplateID: info.MessageTemplateId,
		Phone:             cast.ToInt64(phoneNumber),
		MsgContent:        content.ReplaceContent,
		Status:            10,
		SendDate:          cast.ToInt32(time.Now().Format(timeHelper.DateYMD)),
		Created:           cast.ToInt32(time.Now().Unix()),
		BizId:             *response.Body.BizId,
		SendChannel:       "aliyun",
	}
	return insert
}
