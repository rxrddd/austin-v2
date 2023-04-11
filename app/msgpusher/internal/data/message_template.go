package data

import (
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/utils/cacheHepler"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type IMessageTemplateRepo interface {
	One(ctx context.Context, id int64) (item model.MessageTemplate, err error)
}

type MessageTemplateRepo struct {
	data  *Data
	log   *log.Helper
	cache *cacheHepler.Cache
}

func NewMessageTemplateRepo(data *Data, logger log.Logger) IMessageTemplateRepo {
	return &MessageTemplateRepo{
		data:  data,
		log:   log.NewHelper(log.With(logger, "module", "data/message-template-repo")),
		cache: cacheHepler.NewCache(data.rds, cacheHepler.WithErr(gorm.ErrRecordNotFound)),
	}
}

func (a *MessageTemplateRepo) One(ctx context.Context, id int64) (item model.MessageTemplate, err error) {
	key := fmt.Sprintf("messagetemplate_%d", id)
	err = a.cache.GetOrSet(ctx, key, &item, func(ctx context.Context, v interface{}) error {
		return a.data.db.WithContext(ctx).Where("id", id).First(&item).Error
	})
	return item, err
}
