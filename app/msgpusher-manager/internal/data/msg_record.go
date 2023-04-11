package data

import (
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/dal/model"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/gormHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
)

type IMsgRecordRepo interface {
	GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (items []*model.MsgRecord, total int32, err error)
}

type msgRecordRepo struct {
	data *Data
	log  *log.Helper
}

func NewMsgRecordRepo(data *Data, logger log.Logger) IMsgRecordRepo {
	return &msgRecordRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/msg-record-repo")),
	}
}

func (r *msgRecordRepo) GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (items []*model.MsgRecord, total int32, err error) {
	u := r.data.Query(ctx).MsgRecord
	qx := u.
		Order(u.ID.Desc())
	if emptyHelper.IsNotEmpty(req.RequestId) {
		qx = qx.Where(u.RequestID.Eq(req.RequestId))
	}
	if emptyHelper.IsNotEmpty(req.TemplateId) {
		qx = qx.Where(u.MessageTemplateID.Eq(cast.ToInt64(req.TemplateId)))
	}
	if emptyHelper.IsNotEmpty(req.Channel) {
		qx = qx.Where(u.Channel.Eq(req.Channel))
	}
	items = make([]*model.MsgRecord, 0)
	count, err := qx.Count()
	if err != nil || count <= 0 {
		return items, 0, err
	}
	total = cast.ToInt32(count)
	items, err = qx.Scopes(gormHelper.QueryPage(req.PageNo, req.PageSize)).
		Find()
	return items, total, err
}
