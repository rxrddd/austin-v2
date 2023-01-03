package biz

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/process"
	"austin-v2/pkg/types"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	//"google.golang.org/protobuf/types/known/structpb"
)

type MsgPusherUseCase struct {
	log *log.Helper
	pr  *process.BusinessProcess
	uc  *MessageTemplateUseCase
}

func NewMsgPusherUseCase(
	logger log.Logger,
	pr *process.BusinessProcess,
	uc *MessageTemplateUseCase,
) *MsgPusherUseCase {
	return &MsgPusherUseCase{
		log: log.NewHelper(log.With(logger, "module", "msgpusher-worker/biz/msg-pusher-use-case")),
		pr:  pr,
		uc:  uc,
	}
}
func (s *MsgPusherUseCase) Send(ctx context.Context, in *pb.SendRequest) (resp *pb.SendResponse, err error) {
	if in.MessageParam == nil {
		return nil, errors.Wrapf(errors.New("客户端参数错误1"), "in:%v", in)
	}

	var sendTaskModel = &types.SendTaskModel{
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList: []types.MessageParam{
			{
				Receiver:  in.MessageParam.Receiver,
				Variables: in.MessageParam.Variables.AsMap(),
				Extra:     in.MessageParam.Extra.AsMap(),
			},
		},
	}
	messageTemplate, err := s.uc.One(ctx, sendTaskModel.MessageTemplateId)
	if err != nil {
		return nil, fmt.Errorf("查询模板异常 err:%v 模板id:%d", err, sendTaskModel.MessageTemplateId)
	}

	if err = s.pr.Process(ctx, sendTaskModel, messageTemplate); err != nil {
		return nil, err
	}
	return &pb.SendResponse{
		Code: "send",
	}, nil
}
func (s *MsgPusherUseCase) BatchSend(ctx context.Context, req *pb.BatchSendRequest) (*pb.SendResponse, error) {
	return &pb.SendResponse{
		Code: "send",
	}, nil
}
