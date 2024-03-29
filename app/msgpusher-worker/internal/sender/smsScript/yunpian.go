package smsScript

import (
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/common/dal/model"
	"austin-v2/common/domain/account"
	"austin-v2/common/domain/content_model"
	"austin-v2/pkg/types"
	"austin-v2/utils/accountHelper"
	"austin-v2/utils/contentHelper"
	"austin-v2/utils/jsonHelper"
	"austin-v2/utils/stringHelper"
	"austin-v2/utils/timeHelper"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type YunPian struct {
	logger *log.Helper
	cli    *asynq.Client
	sc     *biz.SendAccountUseCase
}

func NewYunPin(
	logger log.Logger,
	cli *asynq.Client,
	sc *biz.SendAccountUseCase,
) *YunPian {
	return &YunPian{
		logger: log.NewHelper(log.With(logger, "module", "sender/smsScript/yunpian")),
		cli:    cli,
		sc:     sc,
	}
}
func (h *YunPian) Name() string {
	return "yunpian"
}

const sendSmsUrl = "https://sms.yunpian.com/v2/sms/single_send.json"

func (h *YunPian) Send(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	fmt.Println("yunpian sender " + timeHelper.CurrentTimeYMDHIS())

	var acc account.YunPianAccount
	err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc)
	if err != nil {
		return errors.Wrap(err, "yunpian get account err")
	}
	var content content_model.SmsContentModel

	contentHelper.GetContentModel(taskInfo.ContentModel, &content)

	callbackUrl := acc.CallbackUrl
	records := make([]model.SmsRecord, 0)
	for _, receiver := range taskInfo.Receiver {
		// 智能模板发送短信url
		text := content.ReplaceContent

		data := url.Values{
			"apikey":       {acc.ApiKey},
			"mobile":       {receiver},
			"text":         {text},
			"callback_url": {callbackUrl},
		}

		resp, err := httpsPostForm(sendSmsUrl, data)
		if err != nil {
			h.logger.WithContext(ctx).Errorw(
				"msg", "yun pian send msg err",
				"err", err,
				"receiver", receiver,
				"text", text)
			continue
		}
		if resp.Code != 0 {
			h.logger.WithContext(ctx).Errorw(
				"msg", "yun pian send msg code err",
				"error msg", resp.Msg)
			continue
		}

		records = append(records, h.smsRecord(resp, taskInfo, receiver, content))

	}
	_ = ants.Submit(func() {
		if _, err = h.cli.EnqueueContext(ctx, asynq.NewTask("sms.record", jsonHelper.MustToByte(records))); err != nil {
			h.logger.WithContext(ctx).Errorw("msg", "yun pian send publish err", "err", err)
		}
	})

	return nil
}

func httpsPostForm(url string, data url.Values) (*YunPianResp, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respData YunPianResp
	err = json.Unmarshal(body, &respData)
	return &respData, err
}

type YunPianResp struct {
	Code   int     `json:"code"` //0 代表发送成功，其他 code 代表出错，
	Msg    string  `json:"msg"`
	Count  int     `json:"count"` //发送成功短信的计费条数(计费条数：70 个字一条，超出 70 个字时按每 67 字一条计费)
	Fee    float64 `json:"fee"`   //扣费金额，单位：元，类型：双精度浮点型/double
	Unit   string  `json:"unit"`  //计费单位；例如：“RMB”
	Mobile string  `json:"mobile"`
	Sid    int64   `json:"sid"` //短信 id，64 位整型
}

func (h *YunPian) smsRecord(response *YunPianResp,
	info *types.TaskInfo,
	phoneNumber string,
	content content_model.SmsContentModel,
) model.SmsRecord {
	var insert = model.SmsRecord{
		ID:                stringHelper.NextID(),
		MessageTemplateID: info.MessageTemplateId,
		RequestID:         info.RequestId,
		Phone:             cast.ToInt64(phoneNumber),
		MsgContent:        content.ReplaceContent,
		Status:            10,
		ChargingNum:       cast.ToInt32(response.Count),
		SendDate:          cast.ToInt32(time.Now().Format(timeHelper.DateYMD)),
		CreateAt:          time.Now().Unix(),
		BizID:             cast.ToString(response.Sid),
		SendChannel:       "aliyun",
	}
	return insert
}
