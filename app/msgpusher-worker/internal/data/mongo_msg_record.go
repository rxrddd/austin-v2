package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMsgRecordRepo interface {
	InsertMany(ctx context.Context, items []interface{}) error
}

type msgRecordRepo struct {
	data       *Data
	log        *log.Helper
	collection *mongo.Collection
}

func NewMsgRecordRepo(data *Data, logger log.Logger) IMsgRecordRepo {
	return &msgRecordRepo{
		data:       data,
		log:        log.NewHelper(log.With(logger, "module", "data/msg-record-repo")),
		collection: data.mongo.Database("test").Collection("msg_record"),
	}
}

func (r *msgRecordRepo) InsertMany(ctx context.Context, items []interface{}) error {
	_, err := r.collection.InsertMany(ctx, items)
	return err
}
