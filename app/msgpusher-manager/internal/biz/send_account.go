package biz

import (
	pb "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/app/msgpusher-common/model"
	"austin-v2/app/msgpusher-manager/internal/data"
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

func (s *SendAccountUseCase) SendAccountEdit(ctx context.Context, req *pb.SendAccountEditRequest) (*emptypb.Empty, error) {
	var err error
	if req.ID > 0 {
		err = s.repo.SendAccountEdit(ctx, &model.SendAccount{
			ID:          req.ID,
			SendChannel: req.SendChannel,
			Config:      req.Config,
			Title:       req.Title,
		})
	} else {
		err = s.repo.SendAccountCreate(ctx, &model.SendAccount{
			SendChannel: req.SendChannel,
			Config:      req.Config,
			Title:       req.Title,
		})
	}

	return &emptypb.Empty{}, err
}
func (s *SendAccountUseCase) SendAccountChangeStatus(ctx context.Context, req *pb.SendAccountChangeStatusRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.repo.SendAccountChangeStatus(ctx, req.ID, int(req.Status))
}
func (s *SendAccountUseCase) SendAccountList(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountListResp, error) {
	items, total, err := s.repo.SendAccountList(ctx, data.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
		Page:        req.Page,
		PageSize:    req.PageSize,
	})
	rows := make([]*pb.SendAccountRow, 0)
	for _, item := range items {
		rows = append(rows, &pb.SendAccountRow{
			ID:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}

	return &pb.SendAccountListResp{
		Rows:  rows,
		Total: total,
	}, err
}
func (s *SendAccountUseCase) SendAccountQuery(ctx context.Context, req *pb.SendAccountListRequest) (*pb.SendAccountQueryResp, error) {
	items, err := s.repo.SendAccountQuery(ctx, data.SendAccountListRequest{
		Title:       req.Title,
		SendChannel: req.SendChannel,
	})
	rows := make([]*pb.SendAccountRow, 0)
	for _, item := range items {
		rows = append(rows, &pb.SendAccountRow{
			ID:          item.ID,
			Title:       item.Title,
			Config:      item.Config,
			SendChannel: item.SendChannel,
		})
	}

	return &pb.SendAccountQueryResp{
		Rows: rows,
	}, err
}
