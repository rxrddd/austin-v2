package data

import (
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/model"
	"austin-v2/utils/cacheHepler"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/gromHelper"
	"austin-v2/utils/stringHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type IMessageTemplateRepo interface {
	TemplateEdit(ctx context.Context, req *model.MessageTemplate) error
	TemplateCreate(ctx context.Context, req *model.MessageTemplate) error
	TemplateChangeStatus(ctx context.Context, id int64, status int) error
	TemplateList(ctx context.Context, req *domain.TemplateListRequest) (items []model.MessageTemplate, total int64, err error)
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

type TemplateListRequest struct {
	Name        string
	SendChannel string
	Page        int64
	PageSize    int64
}

func (s *messageTemplateRepo) TemplateEdit(ctx context.Context, req *model.MessageTemplate) error {
	return s.data.db.WithContext(ctx).Model(model.MessageTemplate{}).Where("id = ?", req.ID).Updates(req).Error
}
func (s *messageTemplateRepo) TemplateCreate(ctx context.Context, req *model.MessageTemplate) error {
	req.ID = stringHelper.NextID()
	return s.data.db.WithContext(ctx).Model(model.MessageTemplate{}).Create(req).Error
}
func (s *messageTemplateRepo) One(ctx context.Context, id int64) (item model.MessageTemplate, err error) {
	err = s.data.db.WithContext(ctx).Model(model.MessageTemplate{}).Limit(1).Where("id=?", id).Find(&item).Error
	return item, err
}
func (s *messageTemplateRepo) TemplateChangeStatus(ctx context.Context, id int64, status int) error {
	return s.data.db.WithContext(ctx).Model(model.MessageTemplate{}).Where("id = ?", id).UpdateColumn("is_deleted", status).Error
}
func (s *messageTemplateRepo) TemplateList(ctx context.Context, req *domain.TemplateListRequest) (items []model.MessageTemplate, total int64, err error) {
	items = make([]model.MessageTemplate, 0)
	query := s.data.db.WithContext(ctx).Model(items).
		Where("is_deleted=0")
	if emptyHelper.IsNotEmpty(req.Name) {
		query.Where("name like ?", "%"+req.Name+"%")
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		query.Where("send_channel = ?", req.SendChannel)
	}

	query.Count(&total).
		Scopes(gromHelper.Page(req.Page, req.PageSize)).
		Order("id desc").
		Preload("SendAccountItem").
		Find(&items)
	return items, total, nil
}
