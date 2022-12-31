package data

import (
	"austin-v2/app/msgpusher-worker/internal/data/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type IMessageTemplateRepo interface {
	One(ctx context.Context, id int64) (item *model.MessageTemplate, err error)
}

type MessageTemplateRepo struct {
	data *Data
	log  *log.Helper
}

func NewMessageTemplateRepo(data *Data, logger log.Logger) IMessageTemplateRepo {
	return &MessageTemplateRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/message-template-repo")),
	}
}

func (a *MessageTemplateRepo) One(ctx context.Context, id int64) (item *model.MessageTemplate, err error) {
	//key := fmt.Sprintf("messagetemplate_%d", id)
	err = a.data.db.WithContext(ctx).Where("id", id).Limit(1).Find(item).Error
	return item, err
}
