package biz

import (
	"austin-v2/app/msgpusher-manager/internal/data"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/utils/timeHelper"
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

func (s *MsgRecordUseCase) GetMsgRecord(ctx context.Context, req *domain.MsgRecordRequest) (*domain.MsgRecordResp, error) {
	records, total, err := s.repo.GetMsgRecord(ctx, &domain.MsgRecordRequest{
		TemplateId: req.TemplateId,
		RequestId:  req.RequestId,
		Channel:    req.Channel,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var items []*domain.MsgRecordRow
	for _, v := range records {
		items = append(items, &domain.MsgRecordRow{
			MessageTemplateId: v.MessageTemplateID,
			RequestId:         v.RequestID,
			Receiver:          v.Receiver,
			MsgId:             v.MsgID,
			Channel:           v.Channel,
			Msg:               v.Msg,
			SendAt:            v.SendAt,
			CreateAt:          v.CreateAt.Format(timeHelper.DateDefaultLayout),
			SendSinceTime:     v.SendSinceTime,
			ID:                v.ID,
		})
	}
	return &domain.MsgRecordResp{
		Rows:  items,
		Total: total,
	}, nil
}
