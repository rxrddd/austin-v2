package biz

import (
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-manager/internal/data"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/pkg/utils/timeHelper"
	"context"
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

func (s *MessageTemplateUseCase) TemplateEdit(ctx context.Context, req *domain.TemplateEditRequest) (*emptypb.Empty, error) {
	var err error
	var m = &model.MessageTemplate{
		Name:                req.Name,
		IDType:              req.IDType,
		SendChannel:         req.SendChannel,
		TemplateType:        req.TemplateType,
		TemplateSn:          req.TemplateSn,
		MsgType:             req.MsgType,
		ShieldType:          req.ShieldType,
		MsgContent:          req.MsgContent,
		SendAccount:         req.SendAccount,
		SmsChannel:          req.SmsChannel,
		Updated:             time.Now().Unix(),
		DeduplicationConfig: req.DeduplicationConfig,
	}
	if req.ID > 0 {
		m.ID = req.ID
		err = s.repo.TemplateEdit(ctx, m)
	} else {
		err = s.repo.TemplateCreate(ctx, m)
	}
	return &emptypb.Empty{}, err
}
func (s *MessageTemplateUseCase) TemplateChangeStatus(ctx context.Context, id int64, status int) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.repo.TemplateChangeStatus(ctx, id, status)
}

func (s *MessageTemplateUseCase) TemplateList(ctx context.Context, req *domain.TemplateListRequest) (*domain.TemplateListResp, error) {
	items, total, err := s.repo.TemplateList(ctx, req)
	if err != nil {
		return nil, err
	}
	rows := make([]*domain.TemplateListRow, 0)

	for _, item := range items {
		rows = append(rows, &domain.TemplateListRow{
			ID:                  item.ID,
			Name:                item.Name,
			IdType:              int64(item.IDType),
			SendChannel:         int64(item.SendChannel),
			TemplateType:        int64(item.TemplateType),
			MsgType:             int64(item.MsgType),
			ShieldType:          int64(item.ShieldType),
			MsgContent:          item.MsgContent,
			SendAccount:         item.SendAccount,
			SendAccountName:     item.SendAccountItem.Title,
			TemplateSn:          item.TemplateSn,
			SmsChannel:          item.SmsChannel,
			CreateAt:            timeHelper.FormatTimeInt64YMDHIS(item.Created),
			DeduplicationConfig: item.DeduplicationConfig,
		})
	}

	return &domain.TemplateListResp{
		Rows:  rows,
		Total: total,
	}, err
}

func (s *MessageTemplateUseCase) TemplateOne(ctx context.Context, req *domain.TemplateOneRequest) (*domain.TemplateOneResp, error) {
	item, err := s.repo.One(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &domain.TemplateOneResp{
		ID:                  item.ID,
		Name:                item.Name,
		AuditStatus:         item.AuditStatus,
		IDType:              item.IDType,
		SendChannel:         item.SendChannel,
		TemplateType:        item.TemplateType,
		TemplateSn:          item.TemplateSn,
		MsgType:             item.MsgType,
		ShieldType:          item.ShieldType,
		MsgContent:          item.MsgContent,
		SendAccount:         item.SendAccount,
		Creator:             item.Creator,
		Updator:             item.Updator,
		Auditor:             item.Auditor,
		Team:                item.Team,
		Proposer:            item.Proposer,
		SmsChannel:          item.SmsChannel,
		IsDeleted:           item.IsDeleted,
		Created:             item.Created,
		Updated:             item.Updated,
		DeduplicationConfig: item.DeduplicationConfig,
	}, err
}

func (s *MessageTemplateUseCase) TemplateRemove(ctx context.Context, id int64) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.repo.TemplateChangeStatus(ctx, id, 1)
}
