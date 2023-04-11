package data

import (
	"austin-v2/common/dal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ISmsRecordRepo interface {
	Create(ctx context.Context, items []*model.SmsRecord) error
}

type smsRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewSmsRecordRepo(data *Data, logger log.Logger) ISmsRecordRepo {
	return &smsRecordRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/send_account")),
	}
}

func (s *smsRecordRepo) Create(ctx context.Context, items []*model.SmsRecord) error {
	return s.data.db.WithContext(ctx).Model(items).CreateInBatches(items, 500).Error
}
