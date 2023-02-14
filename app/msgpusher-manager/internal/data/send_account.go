package data

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/pkg/utils/cacheHepler"
	"austin-v2/pkg/utils/emptyHelper"
	"austin-v2/pkg/utils/gromHelper"
	"austin-v2/pkg/utils/stringHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type ISendAccountRepo interface {
	SendAccountEdit(ctx context.Context, req *model.SendAccount) error
	SendAccountCreate(ctx context.Context, req *model.SendAccount) error
	SendAccountChangeStatus(ctx context.Context, id int64, status int) error
	SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, total int64, err error)
	SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, err error)
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

func (r *sendAccountRepo) SendAccountEdit(ctx context.Context, req *model.SendAccount) error {
	return r.data.db.WithContext(ctx).Model(model.SendAccount{}).Where("id = ?", req.ID).Updates(req).Error
}
func (r *sendAccountRepo) SendAccountCreate(ctx context.Context, req *model.SendAccount) error {
	req.ID = stringHelper.NextID()
	return r.data.db.WithContext(ctx).Model(model.SendAccount{}).Create(req).Error
}
func (r *sendAccountRepo) SendAccountChangeStatus(ctx context.Context, id int64, status int) error {
	return r.data.db.WithContext(ctx).Model(model.SendAccount{}).Where("id = ?", id).UpdateColumn("status", status).Error
}
func (r *sendAccountRepo) SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, total int64, err error) {
	items = make([]model.SendAccount, 0)
	query := r.data.db.WithContext(ctx).Model(items)
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
func (r *sendAccountRepo) SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (items []model.SendAccount, err error) {
	items = make([]model.SendAccount, 0)
	query := r.data.db.WithContext(ctx).Model(items)
	if emptyHelper.IsNotEmpty(req.Title) {
		query.Where("title like ?", "%"+req.Title+"%")
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		query.Where("send_channel = ?", req.SendChannel)
	}
	query.Where("status = ?", 0).Find(&items)
	return items, nil
}

func (r *sendAccountRepo) One(ctx context.Context, id int64) (item model.SendAccount, err error) {
	key := fmt.Sprintf("sendaccount_%d", id)
	err = r.cache.GetOrSet(ctx, key, &item, func(ctx context.Context, v interface{}) error {
		return r.data.db.WithContext(ctx).Where("id", id).First(&v).Error
	})
	return item, err
}
