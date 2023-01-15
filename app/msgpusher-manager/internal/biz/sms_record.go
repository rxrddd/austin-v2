package biz

import (
	pb "austin-v2/api/msgpusher-manager/v1"
	"austin-v2/app/msgpusher-manager/internal/data"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SmsRecordUseCase struct {
	repo data.ISmsRecordRepo
	log  *log.Helper
}

func NewSmsRecordUseCase(repo data.ISmsRecordRepo, logger log.Logger) *SmsRecordUseCase {
	return &SmsRecordUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/send-account-usecase")),
	}
}

func (s *SmsRecordUseCase) GetSmsRecord(ctx context.Context, req *pb.SmsRecordRequest) (*pb.SmsRecordResp, error) {
	var (
		items       = make([]*pb.SmsRecordRow, 0)
		total int64 = 0
	)

	return &pb.SmsRecordResp{
		Rows:  items,
		Total: total,
	}, nil
}
