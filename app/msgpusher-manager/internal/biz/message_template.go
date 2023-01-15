package biz

import (
	pb "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-manager/internal/data"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type MessageTemplateUseCase struct {
	repo            data.IMessageTemplateRepo
	sendAccountRepo data.ISendAccountRepo
	log             *log.Helper
}

func NewMessageTemplateUseCase(
	repo data.IMessageTemplateRepo,
	logger log.Logger,
) *MessageTemplateUseCase {
	return &MessageTemplateUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/message-template-usecase")),
	}
}

func (s *MessageTemplateUseCase) TemplateEdit(ctx context.Context, req *pb.TemplateEditRequest) (*emptypb.Empty, error) {
	var err error
	if req.ID > 0 {
		err = s.repo.TemplateEdit(ctx, &model.MessageTemplate{
			ID:                  req.ID,
			Name:                req.Name,
			IDType:              int(req.IdType),
			SendChannel:         int(req.SendChannel),
			TemplateType:        int(req.TemplateType),
			TemplateSn:          req.TemplateSn,
			MsgType:             int(req.MsgType),
			ShieldType:          int(req.ShieldType),
			MsgContent:          req.MsgContent,
			SendAccount:         req.SendAccount,
			SmsChannel:          req.SmsChannel,
			Updated:             time.Now().Unix(),
			DeduplicationConfig: req.DeduplicationConfig,
		})
	} else {
		err = s.repo.TemplateCreate(ctx, &model.MessageTemplate{
			Name:                req.Name,
			IDType:              int(req.IdType),
			SendChannel:         int(req.SendChannel),
			TemplateType:        int(req.TemplateType),
			TemplateSn:          req.TemplateSn,
			MsgType:             int(req.MsgType),
			ShieldType:          int(req.ShieldType),
			MsgContent:          req.MsgContent,
			SendAccount:         req.SendAccount,
			SmsChannel:          req.SmsChannel,
			Updated:             time.Now().Unix(),
			Created:             time.Now().Unix(),
			DeduplicationConfig: req.DeduplicationConfig,
		})
	}

	return &emptypb.Empty{}, err
}
func (s *MessageTemplateUseCase) TemplateChangeStatus(ctx context.Context, req *pb.TemplateChangeStatusRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.repo.TemplateChangeStatus(ctx, req.ID, int(req.Status))
}

func (s *MessageTemplateUseCase) TemplateList(ctx context.Context, req *pb.TemplateListRequest) (*pb.TemplateListResp, error) {
	fmt.Println(`TemplateList`)
	items, total, err := s.repo.TemplateList(ctx, data.TemplateListRequest{
		Name:        req.Name,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*pb.TemplateListRow, 0)

	for _, item := range items {
		rows = append(rows, &pb.TemplateListRow{
			ID:              item.ID,
			Name:            item.Name,
			IdType:          int64(item.IDType),
			SendChannel:     int64(item.SendChannel),
			TemplateType:    int64(item.TemplateType),
			MsgType:         int64(item.MsgType),
			ShieldType:      int64(item.ShieldType),
			MsgContent:      item.MsgContent,
			SendAccount:     item.SendAccount,
			SendAccountName: item.SendAccountItem.Title,
			TemplateSn:      item.TemplateSn,
			SmsChannel:      item.SmsChannel,
		})
	}

	return &pb.TemplateListResp{
		Rows:  rows,
		Total: total,
	}, err
}
func (s *MessageTemplateUseCase) TemplateRemove(ctx context.Context, req *pb.TemplateRemoveRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
