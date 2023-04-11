package biz

import (
	"austin-v2/app/msgpusher-worker/internal/data"
	"austin-v2/common/dal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SendAccountUseCase struct {
	repo data.ISendAccountRepo
	log  *log.Helper
}

func NewSendAccountUseCase(repo data.ISendAccountRepo, logger log.Logger) *SendAccountUseCase {
	return &SendAccountUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/send-account-usecase")),
	}
}

func (a *SendAccountUseCase) One(ctx context.Context, id int64) (item model.SendAccount, err error) {
	return a.repo.One(ctx, id)
}
