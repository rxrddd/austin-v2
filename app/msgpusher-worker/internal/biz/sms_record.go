package biz

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-worker/internal/data"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SmsRecordUseCase struct {
	repo data.ISmsRecordRepo
	log  *log.Helper
}

func NewSmsRecordUseCase(repo data.ISmsRecordRepo, logger log.Logger) *SmsRecordUseCase {
	return &SmsRecordUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/send-account-usecase")),
	}
}

func (s *SmsRecordUseCase) Create(ctx context.Context, items []*model.SmsRecord) error {
	return s.repo.Create(ctx, items)
}
