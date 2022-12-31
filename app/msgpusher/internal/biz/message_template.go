package biz

import (
	"austin-v2/app/msgpusher/internal/data"
	"austin-v2/app/msgpusher/internal/data/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MessageTemplateUseCase struct {
	repo data.IMessageTemplateRepo
	log  *log.Helper
}

func NewMessageTemplateUseCase(repo data.IMessageTemplateRepo, logger log.Logger) *MessageTemplateUseCase {
	return &MessageTemplateUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/message-template-usecase")),
	}
}

func (a *MessageTemplateUseCase) One(ctx context.Context, id int64) (item *model.MessageTemplate, err error) {
	return a.repo.One(ctx, id)
}
