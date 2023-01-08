package handler

import (
	"austin-v2/app/msgpusher-common/enums/channelType"
	"austin-v2/app/msgpusher-worker/internal/biz"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/timeHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"golang.org/x/time/rate"
	"time"
)

type EmailHandler struct {
	BaseHandler

	logger  *log.Helper
	rds     redis.Cmdable
	sc      *biz.SendAccountUseCase
	limiter *rate.Limiter
}

func NewEmailHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
) *EmailHandler {
	return &EmailHandler{
		logger:  log.NewHelper(log.With(logger, "module", "sender/sms")),
		rds:     rds,
		sc:      sc,
		limiter: rate.NewLimiter(rate.Every(time.Second*3), 1),
	}
}

func (h *EmailHandler) Name() string {
	return channelType.TypeCodeEn[channelType.Email]
}

func (h *EmailHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	fmt.Println("email sender " + timeHelper.CurrentTimeYMDHIS())
	return nil
	//var content content_model.EmailContentModel
	//contentHelper.GetContentModel(taskInfo.ContentModel, &content)
	//m := gomail.NewMessage()
	//
	//var acc account.EmailAccount
	//err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc)
	//if err != nil {
	//	return errors.Wrap(err, "emailHandler get account err")
	//}
	//
	//m.SetHeader("From", m.FormatAddress(acc.Username, "官方"))
	//
	//m.SetHeader("To", taskInfo.Receiver...) //主送
	//
	//m.SetHeader("Subject", content.Title)
	////发送html格式邮件。
	//m.SetBody("text/html", content.Content)
	//
	//d := gomail.NewDialer(acc.Host, acc.Port, acc.Username, acc.Password)
	//if err := d.DialAndSend(m); err != nil {
	//	return errors.Wrap(err, "emailHandler DialAndSend err")
	//}
	//return nil
}

func (h *EmailHandler) Allow(ctx context.Context, _ *types.TaskInfo) bool {
	return h.limiter.AllowN(time.Now(), 1)
}
