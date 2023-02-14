package data

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/pkg/utils/cacheHepler"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type IMessageTemplateRepo interface {
	One(ctx context.Context, id int64) (item model.MessageTemplate, err error)
}

type messageTemplateRepo struct {
	data  *Data
	log   *log.Helper
	cache *cacheHepler.Cache
}

func NewMessageTemplateRepo(data *Data, logger log.Logger) IMessageTemplateRepo {
	return &messageTemplateRepo{
		data:  data,
		log:   log.NewHelper(log.With(logger, "module", "data/message-template-repo")),
		cache: cacheHepler.NewCache(data.rds),
	}
}

func (a *messageTemplateRepo) One(ctx context.Context, id int64) (item model.MessageTemplate, err error) {
	key := fmt.Sprintf("messagetemplate_%d", id)
	err = a.cache.GetOrSet(ctx, key, &item, func(ctx context.Context, v interface{}) error {
		return a.data.db.WithContext(ctx).Where("id", id).First(&v).Error
	})
	return item, err
}
