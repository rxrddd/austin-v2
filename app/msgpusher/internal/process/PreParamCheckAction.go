package process

import (
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/pkg/types"
	"context"
	"github.com/pkg/errors"
)

type PreParamCheckAction struct {
}

func NewPreParamCheckAction() *PreParamCheckAction {
	return &PreParamCheckAction{}
}

func (p *PreParamCheckAction) Process(_ context.Context, sendTaskModel *types.SendTaskModel, messageTemplate model.MessageTemplate) error {

	if sendTaskModel.MessageTemplateId == 0 || len(sendTaskModel.MessageParamList) <= 0 {
		return errors.Wrapf(clientParamsErr, "PreParamCheckAction sendTaskModel:%v", sendTaskModel)
	}
	// 过滤 receiver=null 的messageParam
	var newRows []types.MessageParam
	for _, param := range sendTaskModel.MessageParamList {
		if param.Receiver != "" {
			newRows = append(newRows, param)
		}
	}
	if len(newRows) <= 0 {
		return errors.Wrapf(clientParamsErr, "PreParamCheckAction 过滤结果参数为空 sendTaskModel:%v", sendTaskModel)
	}
	sendTaskModel.MessageParamList = newRows
	return nil
}
