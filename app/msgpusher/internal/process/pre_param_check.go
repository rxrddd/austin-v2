package process

import (
	pb "austin-v2/api/msgpusher/v1"
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/pkg/types"
	"context"
)

type PreParamCheckAction struct {
}

func NewPreParamCheckAction() *PreParamCheckAction {
	return &PreParamCheckAction{}
}

func (p *PreParamCheckAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel, _ model.MessageTemplate) error {
	if sendTaskModel.MessageTemplateId == 0 || len(sendTaskModel.MessageParamList) <= 0 {
		return pb.ErrorClientParamsError("客户端参数错误")
	}
	// 过滤 receiver=null 的messageParam
	var newRows []types.MessageParam
	for _, param := range sendTaskModel.MessageParamList {
		if param.Receiver != "" {
			newRows = append(newRows, param)
		}
	}
	if len(newRows) <= 0 {
		return pb.ErrorPreParamAllFilter("无可发送消息")
	}
	sendTaskModel.MessageParamList = newRows
	return nil
}
