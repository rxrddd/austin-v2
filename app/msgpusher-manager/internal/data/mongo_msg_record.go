package data

import (
	"austin-v2/app/msgpusher-common/model/mongo_model"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/pkg/utils/emptyHelper"
	"austin-v2/pkg/utils/jsonHelper"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMsgRecordRepo interface {
	GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (items []*mongo_model.MsgRecord, total int64, err error)
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

func (r *msgRecordRepo) GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (items []*mongo_model.MsgRecord, total int64, err error) {
	query := bson.M{}

	if emptyHelper.IsNotEmpty(req.RequestId) {
		query["request_id"] = req.RequestId
	}
	if emptyHelper.IsNotEmpty(req.TemplateId) {
		query["message_template_id"] = cast.ToInt64(req.TemplateId)
	}
	if emptyHelper.IsNotEmpty(req.Channel) {
		query["channel"] = req.Channel
	}
	fmt.Println(`query`, jsonHelper.MustToString(query))
	opt := options.FindOptions{}
	opt.SetSkip((req.Page - 1) * req.PageSize)
	opt.SetLimit(req.PageSize)
	opt.SetSort(bson.M{"create_at": -1})
	opt.SetProjection(bson.D{
		{"task_info", 0},
	})
	cur, err := r.collection.Find(ctx, query, &opt)
	if err != nil {
		return nil, 0, err
	}
	var resp []*mongo_model.MsgRecord
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		item := &mongo_model.MsgRecord{}
		err := cur.Decode(item)
		if err != nil {
			return nil, 0, err
		}
		resp = append(resp, item)
	}
	copt := options.CountOptions{}
	total, err = r.collection.CountDocuments(ctx, query, &copt)
	if err != nil {
		return nil, 0, err
	}
	return resp, total, nil
}
