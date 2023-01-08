package smsScript

import (
	"austin-v2/app/msgpusher-common/domain/account"
	"austin-v2/app/msgpusher-common/domain/content_model"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/app/msgpusher-worker/internal/data/model"
	"austin-v2/app/msgpusher-worker/internal/pkg/utils/accountHelper"
	"austin-v2/pkg/mq"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/contentHelper"
	"austin-v2/pkg/utils/stringHelper"
	"austin-v2/pkg/utils/timeHelper"
	"context"
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	smsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
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
	err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc)
	if err != nil {
		return errors.Wrap(err, "yunpian get account err")
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
		bytes, _ := json.Marshal(taskInfo.MessageParam.Variables)

		request.SetTemplateParam(string(bytes))
		response, err := cli.SendSms(request)
		if err != nil {
			return fmt.Errorf("Client.Send() error = %v", err)
		}
		if *response.Body.Code == "OK" {
			records = append(records, h.smsRecord(response, taskInfo.MessageTemplateId, receiver, content))
		}
	}
	marshal, _ := json.Marshal(records)
	err = h.mqHelper.Publish(marshal, "sms.record")
	if err != nil {
		h.logger.WithContext(ctx).Errorw("msg", "aliyun send publish err", "err", err)
	}

	return nil
}
func (h *AliyunSms) smsRecord(response *smsapi.SendSmsResponse, messageTemplateId int64, phoneNumber string, content content_model.SmsContentModel) model.SmsRecord {
	requestId := *response.Body.RequestId
	var insert = model.SmsRecord{
		ID:                stringHelper.NextID(),
		MessageTemplateID: messageTemplateId,
		Phone:             cast.ToInt64(phoneNumber),
		MsgContent:        content.ReplaceContent,
		Status:            10,
		SendDate:          cast.ToInt32(time.Now().Format(timeHelper.DateYMD)),
		Created:           cast.ToInt32(time.Now().Unix()),
		RequestId:         requestId,
		BizId:             *response.Body.BizId,
		SendChannel:       "aliyun",
	}
	return insert
}
