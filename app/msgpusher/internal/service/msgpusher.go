package service

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/biz"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
)

type MsgPusherService struct {
	pb.UnimplementedMsgPusherServer
	uc  *biz.GreeterUsecase
	b   broker.Broker
	log *log.Helper
}

func NewMsgPusherService(uc *biz.GreeterUsecase, b broker.Broker, logger log.Logger) *MsgPusherService {
	return &MsgPusherService{
		uc:  uc,
		b:   b,
		log: log.NewHelper(logger),
	}
}

type Msg struct {
	Name string
}

func (s *MsgPusherService) Send(ctx context.Context, req *pb.SendRequest) (*pb.SendResponse, error) {
	msg := Msg{Name: "张三"}
	buf, _ := json.Marshal(&msg)

	err := s.b.Publish("", buf,
		rabbitmq.WithPublishDeclareQueue("test_zhangsan", true, false, map[string]interface{}{}, map[string]interface{}{}),
	)
	if err != nil {
		s.log.Error(`err`, err)
	}
	s.log.Info(`buf`, string(buf))

	return &pb.SendResponse{}, nil
}
func (s *MsgPusherService) BatchSend(ctx context.Context, req *pb.BatchSendRequest) (*pb.SendResponse, error) {
	return &pb.SendResponse{}, nil
}
