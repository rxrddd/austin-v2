package data

import (
	"austin-v2/common/model"
	"austin-v2/utils/cacheHepler"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type ISendAccountRepo interface {
	One(ctx context.Context, id int64) (item model.SendAccount, err error)
}

type sendAccountRepo struct {
	data  *Data
	log   *log.Helper
	cache *cacheHepler.Cache
}

func NewSendAccountRepo(data *Data, logger log.Logger) ISendAccountRepo {
	return &sendAccountRepo{
		data:  data,
		log:   log.NewHelper(log.With(logger, "module", "data/send_account")),
		cache: cacheHepler.NewCache(data.rds),
	}
}

func (a *sendAccountRepo) One(ctx context.Context, id int64) (item model.SendAccount, err error) {
	key := fmt.Sprintf("sendaccount_%d", id)
	err = a.cache.GetOrSet(ctx, key, &item, func(ctx context.Context, v interface{}) error {
		return a.data.db.WithContext(ctx).Where("id", id).First(&v).Error
	})
	return item, err
}
