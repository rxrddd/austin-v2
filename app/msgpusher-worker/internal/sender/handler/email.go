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
	"golang.org/x/time/rate"
	"gopkg.in/gomail.v2"
	"strings"
	"time"
)

type EmailHandler struct {
	BaseHandler

	logger  *log.Helper
	rds     redis.Cmdable
	sc      *biz.SendAccountUseCase
	mrr     data.IMsgRecordRepo
	limiter *rate.Limiter
}

func NewEmailHandler(
	logger log.Logger,
	rds redis.Cmdable,
	sc *biz.SendAccountUseCase,
	mrr data.IMsgRecordRepo,
) *EmailHandler {
	return &EmailHandler{
		logger:  log.NewHelper(log.With(logger, "module", "sender/sms")),
		rds:     rds,
		sc:      sc,
		limiter: rate.NewLimiter(rate.Every(time.Second*3), 1),
		mrr:     mrr,
	}
}

func (h *EmailHandler) Name() string {
	return channelType.TypeCodeEn[channelType.Email]
}

func (h *EmailHandler) Execute(ctx context.Context, taskInfo *types.TaskInfo) (err error) {
	var content content_model.EmailContentModel
	contentHelper.GetContentModel(taskInfo.ContentModel, &content)
	m := gomail.NewMessage()

	var acc account.EmailAccount

	if err = accountHelper.GetAccount(ctx, h.sc, taskInfo.SendAccount, &acc); err != nil {
		h.logger.WithContext(ctx).Errorw(
			"msg", "emailHandler send err",
			"err", err,
			"requestId", taskInfo.RequestId)
		return err
	}

	m.SetHeader("From", m.FormatAddress(acc.Username, "官方"))

	m.SetHeader("To", taskInfo.Receiver...) //主送

	m.SetHeader("Subject", content.Title)
	//发送html格式邮件。
	m.SetBody("text/html", content.Content)

	record := h.getRecord(taskInfo, strings.Join(taskInfo.Receiver, ","))
	record.Channel = h.Name()

	if err := gomail.NewDialer(acc.Host, acc.Port, acc.Username, acc.Password).
		DialAndSend(m); err != nil {
		h.logger.WithContext(ctx).Errorw(
			"msg", "emailHandler send err",
			"err", err,
			"requestId", taskInfo.RequestId)
		record.Msg = "推送失败: " + err.Error()
	} else {
		record.Msg = "推送成功"
		h.logger.WithContext(ctx).Infow(
			"msg", "emailHandler send success",
			"requestId", taskInfo.RequestId)
	}
	_ = ants.Submit(func() {
		_ = h.mrr.InsertMany(ctx, []*model.MsgRecord{record})
	})
	return nil
}

func (h *EmailHandler) Allow(ctx context.Context, _ *types.TaskInfo) bool {
	return h.limiter.AllowN(time.Now(), 1)
}
