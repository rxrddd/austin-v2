package data

import (
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/dal/model"
	"austin-v2/utils/cacheHepler"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/gormHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
)

type ISendAccountRepo interface {
	SendAccountEdit(ctx context.Context, req *model.SendAccount) error
	SendAccountCreate(ctx context.Context, req *model.SendAccount) error
	SendAccountChangeStatus(ctx context.Context, id int32, status int32) error
	SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (items []*model.SendAccount, total int32, err error)
	SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (items []*model.SendAccount, err error)
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
	return r.data.db.WithContext(ctx).Model(model.SendAccount{}).Create(req).Error
}
func (r *sendAccountRepo) SendAccountChangeStatus(ctx context.Context, id int32, status int32) error {
	u := r.data.Query(ctx).SendAccount
	_, err := u.
		Where(u.ID.Eq(id)).
		UpdateSimple(u.Status.Value(status))
	return err
}
func (r *sendAccountRepo) SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (items []*model.SendAccount, total int32, err error) {
	u := r.data.Query(ctx).SendAccount
	qx := u.
		Order(u.ID.Desc())
	if emptyHelper.IsNotEmpty(req.Title) {
		qx = qx.Where(u.Title.Like("%" + req.Title + "%"))
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		qx = qx.Where(u.SendChannel.Eq(req.SendChannel))
	}

	items = make([]*model.SendAccount, 0)
	count, err := qx.Count()
	if err != nil || count <= 0 {
		return items, 0, err
	}
	total = cast.ToInt32(count)
	items, err = qx.Scopes(gormHelper.QueryPage(req.PageNo, req.PageSize)).
		Find()
	return items, total, err
}
func (r *sendAccountRepo) SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (items []*model.SendAccount, err error) {
	u := r.data.Query(ctx).SendAccount
	qx := u.
		Order(u.ID.Desc())
	if emptyHelper.IsNotEmpty(req.Title) {
		qx = qx.Where(u.Title.Like("%" + req.Title + "%"))
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		qx = qx.Where(u.SendChannel.Eq(req.SendChannel))
	}

	items = make([]*model.SendAccount, 0)
	items, err = qx.Scopes(gormHelper.QueryPage(req.PageNo, req.PageSize)).
		Find()
	return items, nil
}

func (r *sendAccountRepo) One(ctx context.Context, id int64) (item model.SendAccount, err error) {
	key := fmt.Sprintf("sendaccount_%d", id)
	err = r.cache.GetOrSet(ctx, key, &item, func(ctx context.Context, v interface{}) error {
		return r.data.db.WithContext(ctx).Where("id", id).First(&v).Error
	})
	return item, err
}
