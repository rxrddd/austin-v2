package data

import (
	"austin-v2/common/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type IMsgRecordRepo interface {
	InsertMany(ctx context.Context, items []*model.MsgRecord) error
}

type msgRecordRepo struct {
	data *Data
	log  *log.Helper
}

// NewMysqlMsgRecordRepo mysql版本的消息记录接口
func NewMysqlMsgRecordRepo(data *Data, logger log.Logger) IMsgRecordRepo {
	return &msgRecordRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/mysql-msg-record-repo")),
	}
}

func (r *msgRecordRepo) InsertMany(ctx context.Context, items []*model.MsgRecord) error {
	return r.data.db.Model(&model.MsgRecord{}).WithContext(ctx).CreateInBatches(items, 100).Error
}
