package service

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MsgPusherService struct {
	pb.UnimplementedMsgPusherServer
	uc  *biz.MsgPusherUseCase
	log *log.Helper
	//pr  *process.BusinessProcess
}

func NewMsgPusherService(
	uc *biz.MsgPusherUseCase,
	logger log.Logger,
	//pr *process.BusinessProcess,
) *MsgPusherService {
	return &MsgPusherService{
		uc:  uc,
		log: log.NewHelper(logger),
		//pr:  pr,
	}
}

type Msg struct {
	Name string
}

func (s *MsgPusherService) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	return s.uc.Send(ctx, req)
}
func (s *MsgPusherService) BatchSend(ctx context.Context, req *pb.BatchSendRequest) (*pb.SendResponse, error) {
	return s.uc.BatchSend(ctx, req)
}
