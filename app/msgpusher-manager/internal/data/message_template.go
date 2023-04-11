package data

import (
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/dal/model"
	"austin-v2/utils/cacheHepler"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/gormHelper"
	"austin-v2/utils/idHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
)

type messageTemplateRepo struct {
	data  *Data
	log   *log.Helper
	cache *cacheHepler.Cache
}
type IMessageTemplateRepo interface {
	TemplateEdit(ctx context.Context, req *model.MessageTemplate) error
	TemplateCreate(ctx context.Context, req *model.MessageTemplate) error
	One(ctx context.Context, id int64) (item *model.MessageTemplate, err error)
	TemplateChangeStatus(ctx context.Context, id int64, status int32) error
	TemplateList(ctx context.Context, req *domain.TemplateListRequest) (m []*model.MessageTemplate, total int32, err error)
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

func (r *messageTemplateRepo) TemplateEdit(ctx context.Context, req *model.MessageTemplate) error {
	q := r.data.Query(ctx).MessageTemplate
	_, err := q.Where(
		q.ID.Eq(req.ID),
	).
		Select(
			q.Name,
			q.AuditStatus,
			q.IDType,
			q.SendChannel,
			q.TemplateType,
			q.MsgType,
			q.ShieldType,
			q.MsgContent,
			q.SendAccount,
			q.DeduplicationConfig,
			q.TemplateSn,
			q.SmsChannel,
		).
		Updates(req)
	return err
}
func (r *messageTemplateRepo) TemplateCreate(ctx context.Context, req *model.MessageTemplate) error {
	req.ID = idHelper.NextID()
	req.Status = 1
	q := r.data.Query(ctx).MessageTemplate
	return q.Create(req)
}
func (r *messageTemplateRepo) One(ctx context.Context, id int64) (item *model.MessageTemplate, err error) {
	q := r.data.Query(ctx).MessageTemplate
	return q.Where(q.ID.Eq(id)).First()
}

func (r *messageTemplateRepo) TemplateChangeStatus(ctx context.Context, id int64, status int32) error {
	q := r.data.Query(ctx).MessageTemplate
	_, err := q.Where(q.ID.Eq(id)).UpdateSimple(q.Status.Value(status))
	return err
}

func (r *messageTemplateRepo) TemplateList(ctx context.Context, req *domain.TemplateListRequest) (m []*model.MessageTemplate, total int32, err error) {
	u := r.data.Query(ctx).MessageTemplate
	qx := u.Where(u.Status.Neq(0)).
		Order(u.ID.Desc())
	if emptyHelper.IsNotEmpty(req.Name) {
		qx = qx.Where(u.Name.Like("%" + req.Name + "%"))
	}
	if emptyHelper.IsNotEmpty(req.SendChannel) {
		qx = qx.Where(u.SendChannel.Eq(cast.ToInt32(req.SendChannel)))
	}
	m = make([]*model.MessageTemplate, 0)
	count, err := qx.Count()
	if err != nil || count <= 0 {
		return m, 0, err
	}
	total = cast.ToInt32(count)
	res, err := qx.Scopes(gormHelper.QueryPage(req.PageNo, req.PageSize)).
		Preload(u.SendAccountItem).
		Find()
	return res, total, err
}
