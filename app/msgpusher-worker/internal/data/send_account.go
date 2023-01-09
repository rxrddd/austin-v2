package data

import (
	"austin-v2/app/msgpusher-common/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ISendAccountRepo interface {
	One(ctx context.Context, id int64) (item *model.SendAccount, err error)
}

type SendAccountRepo struct {
	data *Data
	log  *log.Helper
}

func NewSendAccountRepo(data *Data, logger log.Logger) ISendAccountRepo {
	return &SendAccountRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/send_account")),
	}
}

func (a *SendAccountRepo) One(ctx context.Context, id int64) (item *model.SendAccount, err error) {
	err = a.data.db.WithContext(ctx).Where("id", id).Limit(1).Find(&item).Error
	return item, err
}
