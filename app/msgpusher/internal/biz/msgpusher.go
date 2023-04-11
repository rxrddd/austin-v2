package biz

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/process"
	"austin-v2/pkg/types"
	"austin-v2/utils/stringHelper"
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

//Send 单条推送消息
func (s *MsgPusherUseCase) Send(ctx context.Context, in *pb.SendRequest) (resp *pb.SendResponse, err error) {
	requestId := stringHelper.GenerateUUID()
	if in.MessageParam == nil {
		return nil, pb.ErrorClientParamsError("客户端参数错误")
	}

	var sendTaskModel = &types.SendTaskModel{
		RequestId:         requestId,
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
	if messageTemplate.ID <= 0 || err != nil {
		return nil, pb.ErrorSearchMessageTemplate("查询模板异常 模板id:%d", sendTaskModel.MessageTemplateId)
	}
	if err = s.pr.Process(ctx, sendTaskModel, messageTemplate); err != nil {
		return nil, err
	}

	return &pb.SendResponse{
		RequestId: requestId,
	}, nil
}

//BatchSend 批量推送消息
func (s *MsgPusherUseCase) BatchSend(ctx context.Context, in *pb.BatchSendRequest) (*pb.SendResponse, error) {
	requestId := stringHelper.GenerateUUID()
	if in.MessageParam == nil {
		return nil, pb.ErrorClientParamsError("客户端参数错误")
	}
	messageParamList := make([]types.MessageParam, 0)

	for _, msg := range in.MessageParam {
		messageParamList = append(messageParamList, types.MessageParam{
			Receiver:  msg.Receiver,
			Variables: msg.Variables.AsMap(),
			Extra:     msg.Extra.AsMap(),
		})
	}

	var sendTaskModel = &types.SendTaskModel{
		RequestId:         requestId,
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList:  messageParamList,
	}
	messageTemplate, err := s.uc.One(ctx, sendTaskModel.MessageTemplateId)
	if err != nil {
		return nil, pb.ErrorSearchMessageTemplate("查询模板异常 err:%v 模板id:%d", err, sendTaskModel.MessageTemplateId)
	}

	if err = s.pr.Process(ctx, sendTaskModel, messageTemplate); err != nil {
		return nil, err
	}

	return &pb.SendResponse{
		RequestId: requestId,
	}, nil
}
