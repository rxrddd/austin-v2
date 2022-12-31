package service

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/biz"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

type MsgPusherService struct {
	pb.UnimplementedMsgPusherServer
	uc  *biz.GreeterUsecase
	hs  *biz.HandleUsecase
	b   broker.Broker
	log *log.Helper
}

func NewMsgPusherService(
	uc *biz.GreeterUsecase,
	hs *biz.HandleUsecase,
	b broker.Broker,
	logger log.Logger,
) *MsgPusherService {
	return &MsgPusherService{
		uc:  uc,
		hs:  hs,
		b:   b,
		log: log.NewHelper(logger),
	}
}

type Msg struct {
	Name string
}

func (s *MsgPusherService) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {

	fmt.Println(`send`, s.hs.Handle(ctx, "sms"))
	fmt.Println(`send`, s.hs.Handle(ctx, "sms1"))

	return &pb.SendResponse{}, nil
}
func (s *MsgPusherService) BatchSend(ctx context.Context, req *pb.BatchSendRequest) (*pb.SendResponse, error) {
	return &pb.SendResponse{}, nil
}
