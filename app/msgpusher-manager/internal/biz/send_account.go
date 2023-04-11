package biz

import (
	"austin-v2/app/msgpusher-manager/internal/data"
	"austin-v2/app/msgpusher-manager/internal/domain"
	"austin-v2/common/dal/model"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SendAccountUseCase struct {
	repo data.ISendAccountRepo
	log  *log.Helper
}

func NewSendAccountUseCase(repo data.ISendAccountRepo, logger log.Logger) *SendAccountUseCase {
	return &SendAccountUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/send-account-usecase")),
	}
}

func (s *SendAccountUseCase) SendAccountEdit(ctx context.Context, req *domain.SendAccountEditRequest) (*emptypb.Empty, error) {
	var err error
	m := &model.SendAccount{
		ID:          req.ID,
		SendChannel: req.SendChannel,
		Config:      req.Config,
		Title:       req.Title,
	}
	if req.ID > 0 {
		err = s.repo.SendAccountEdit(ctx, m)
	} else {
		err = s.repo.SendAccountCreate(ctx, m)
	}

	return &emptypb.Empty{}, err
}
func (s *SendAccountUseCase) SendAccountChangeStatus(ctx context.Context, id int32, status int32) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.repo.SendAccountChangeStatus(ctx, id, status)
}
func (s *SendAccountUseCase) SendAccountList(ctx context.Context, req *domain.SendAccountListRequest) (*domain.SendAccountListResp, error) {
	items, total, err := s.repo.SendAccountList(ctx, req)
	rows := make([]*domain.SendAccountRow, 0)
	for _, item := range items {
		rows = append(rows, &domain.SendAccountRow{
			ID:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
			Status:      item.Status,
		})
	}

	return &domain.SendAccountListResp{
		Rows:  rows,
		Total: total,
	}, err
}
func (s *SendAccountUseCase) SendAccountQuery(ctx context.Context, req *domain.SendAccountListRequest) (*domain.SendAccountQueryResp, error) {
	items, err := s.repo.SendAccountQuery(ctx, req)
	rows := make([]*domain.SendAccountRow, 0)
	for _, item := range items {
		rows = append(rows, &domain.SendAccountRow{
			ID:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}

	return &domain.SendAccountQueryResp{
		Rows: rows,
	}, err
}
