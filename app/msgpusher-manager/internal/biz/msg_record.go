package biz

import (
	pb "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/app/msgpusher-manager/internal/data"
	"austin-v2/pkg/utils/timeHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MsgRecordUseCase struct {
	repo data.IMsgRecordRepo
	log  *log.Helper
}

func NewMsgRecordUseCase(repo data.IMsgRecordRepo, logger log.Logger) *MsgRecordUseCase {
	return &MsgRecordUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/msg-record-usecase")),
	}
}

func (s *MsgRecordUseCase) GetMsgRecord(ctx context.Context, req *pb.MsgRecordRequest) (*pb.MsgRecordResp, error) {
	var (
		items       = make([]*pb.MsgRecordRow, 0)
		total int64 = 0
	)
	records, total, err := s.repo.GetMsgRecord(ctx, &data.MsgRecordRequest{
		TemplateId: req.TemplateId,
		RequestId:  req.RequestId,
		Channel:    req.Channel,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range records {
		items = append(items, &pb.MsgRecordRow{
			MessageTemplateId: v.MessageTemplateID,
			RequestId:         v.RequestID,
			Receiver:          v.Receiver,
			MsgId:             v.MsgId,
			Channel:           v.Channel,
			Msg:               v.Msg,
			SendAt:            v.SendAt,
			CreateAt:          v.CreateAt.Format(timeHelper.DateDefaultLayout),
			SendSinceTime:     v.SendSinceTime,
			ID:                v.ID,
		})
	}
	return &pb.MsgRecordResp{
		Rows:  items,
		Total: total,
	}, nil
}
