package biz

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/process"
	"austin-v2/pkg/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type MsgPusherUseCase struct {
	log *log.Helper
	pr  *process.BusinessProcess
}

func NewMsgPusherUseCase(
	logger log.Logger,
	pr *process.BusinessProcess,
) *MsgPusherUseCase {
	return &MsgPusherUseCase{
		log: log.NewHelper(log.With(logger, "module", "msgpusher-worker/biz/msg-pusher-use-case")),
		//pr:pr,
	}
}
func (s *MsgPusherUseCase) Send(ctx context.Context, in *pb.SendRequest) (resp *pb.SendResponse, err error) {
	if in.MessageParam == nil {
		return nil, errors.Wrapf(errors.New("客户端参数错误"), "in:%v", in)
	}
	variables := make(map[string]interface{})
	extra := make(map[string]interface{})
	err = json.Unmarshal([]byte(in.MessageParam.Variables), &variables)
	if err != nil {
		return nil, errors.Wrapf(errors.New("客户端参数错误"), "in:%v", in)
	}
	err = json.Unmarshal([]byte(in.MessageParam.Extra), &extra)
	if err != nil {
		return nil, errors.Wrapf(errors.New("客户端参数解析错误"), "in:%v", in)
	}
	var sendTaskModel = &types.SendTaskModel{
		MessageTemplateId: in.MessageTemplateId,
		MessageParamList: []types.MessageParam{
			{
				Receiver:  in.MessageParam.Receiver,
				Variables: variables,
				Extra:     extra,
			},
		},
	}
	fmt.Println(sendTaskModel)

	if err = s.pr.Process(ctx, sendTaskModel); err != nil {
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
