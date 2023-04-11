package data

import (
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/model"
	"austin-v2/utils/emptyHelper"
	"austin-v2/utils/gromHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
)

type IMsgRecordRepo interface {
	GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (items []*model.MsgRecord, total int64, err error)
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

func (r *msgRecordRepo) GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (items []*model.MsgRecord, total int64, err error) {
	query := r.data.db.WithContext(ctx).Model(&model.MsgRecord{})
	if emptyHelper.IsNotEmpty(req.RequestId) {
		query.Where("request_id = ?", req.RequestId)
	}
	if emptyHelper.IsNotEmpty(req.TemplateId) {
		query.Where("message_template_id = ?", cast.ToInt64(req.TemplateId))
	}
	if emptyHelper.IsNotEmpty(req.Channel) {
		query.Where("channel = ?", req.Channel)
	}
	query.Count(&total).
		Scopes(gromHelper.Page(req.Page, req.PageSize)).
		Find(&items)
	return items, total, nil
}
