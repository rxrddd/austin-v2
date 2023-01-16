package data

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/pkg/utils/cacheHepler"
	"austin-v2/pkg/utils/emptyHelper"
	"austin-v2/pkg/utils/gromHelper"
	"austin-v2/pkg/utils/stringHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ISendAccountRepo interface {
	SendAccountEdit(ctx context.Context, req *model.SendAccount) error
	SendAccountCreate(ctx context.Context, req *model.SendAccount) error
	SendAccountChangeStatus(ctx context.Context, id int64, status int) error
	SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, total int64, err error)
	SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, err error)
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

func (s *sendAccountRepo) SendAccountEdit(ctx context.Context, req *model.SendAccount) error {
	return s.data.db.WithContext(ctx).Model(model.SendAccount{}).Where("id = ?", req.ID).Updates(req).Error
}
func (s *sendAccountRepo) SendAccountCreate(ctx context.Context, req *model.SendAccount) error {
	req.ID = stringHelper.NextID()
	return s.data.db.WithContext(ctx).Model(model.SendAccount{}).Create(req).Error
}
func (s *sendAccountRepo) SendAccountChangeStatus(ctx context.Context, id int64, status int) error {
	return s.data.db.WithContext(ctx).Model(model.SendAccount{}).Where("id = ?", id).UpdateColumn("status", status).Error
}
func (s *sendAccountRepo) SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, total int64, err error) {
	items = make([]model.SendAccount, 0)
	query := s.data.db.WithContext(ctx).Model(items)
	if emptyHelper.IsNotEmpty(req.Title) {
		query.Where("title like ?", "%"+req.Title+"%")
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		query.Where("send_channel = ?", req.SendChannel)
	}
	query.Count(&total).
		Scopes(gromHelper.Page(req.Page, req.PageSize)).
		Find(&items)
	return items, total, nil
}
func (s *sendAccountRepo) SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, err error) {
	items = make([]model.SendAccount, 0)
	query := s.data.db.WithContext(ctx).Model(items)
	if emptyHelper.IsNotEmpty(req.Title) {
		query.Where("title like ?", "%"+req.Title+"%")
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		query.Where("send_channel = ?", req.SendChannel)
	}
	query.Where("status = ?", 0).Find(&items)
	return items, nil
}
