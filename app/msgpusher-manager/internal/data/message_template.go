package data

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/pkg/utils/cacheHepler"
	"austin-v2/pkg/utils/emptyHelper"
	"austin-v2/pkg/utils/gromHelper"
	"austin-v2/pkg/utils/stringHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type IMessageTemplateRepo interface {
	TemplateEdit(ctx context.Context, req *model.MessageTemplate) error
	TemplateCreate(ctx context.Context, req *model.MessageTemplate) error
	TemplateChangeStatus(ctx context.Context, id int64, status int) error
	TemplateList(ctx context.Context, req TemplateListRequest) (items []model.MessageTemplate, total int64, err error)
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

type TemplateListRequest struct {
	Name        string
	SendChannel string
	Page        int64
	PageSize    int64
}

func (s *messageTemplateRepo) TemplateEdit(ctx context.Context, req *model.MessageTemplate) error {
	return s.data.db.WithContext(ctx).Where("id = ?", req.ID).Updates(req).Error
}
func (s *messageTemplateRepo) TemplateCreate(ctx context.Context, req *model.MessageTemplate) error {
	req.ID = stringHelper.NextID()
	return s.data.db.WithContext(ctx).Create(req).Error
}
func (s *messageTemplateRepo) TemplateChangeStatus(ctx context.Context, id int64, status int) error {
	return s.data.db.WithContext(ctx).Where("id = ?", id).UpdateColumn("status", status).Error
}
func (s *messageTemplateRepo) TemplateList(ctx context.Context, req TemplateListRequest) (items []model.MessageTemplate, total int64, err error) {
	items = make([]model.MessageTemplate, 0)
	query := s.data.db.WithContext(ctx).Model(items)
	if emptyHelper.IsNotEmpty(req.Name) {
		query.Where("name like ?", "%"+req.Name+"%")
	}

	query.Count(&total).
		Scopes(gromHelper.Page(req.Page, req.PageSize)).
		Order("id desc").
		Preload("SendAccountItem").
		Find(&items)
	return items, total, nil
}