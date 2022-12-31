package process

import (
	"austin-v2/app/msgpusher-common/domain/content_model"
	"austin-v2/app/msgpusher/internal/biz"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"austin-v2/pkg/utils/transformHelper"
	"context"
	"github.com/pkg/errors"
	"strings"
)

type AssembleAction struct {
	uc *biz.MessageTemplateUseCase
}

func NewAssembleAction(
	uc *biz.MessageTemplateUseCase,
) *AssembleAction {
	return &AssembleAction{
		uc: uc,
	}
}

func (p *AssembleAction) Process(ctx context.Context, sendTaskModel *types.SendTaskModel) error {
	messageParamList := sendTaskModel.MessageParamList

	messageTemplate, err := p.uc.One(ctx, sendTaskModel.MessageTemplateId)
	if err != nil {
		return errors.Wrapf(sendErr, "查询模板异常 err:%v 模板id:%d", err, sendTaskModel.MessageTemplateId)
	}
	contentModel := content_model.GetBuilderContentBySendChannel(messageTemplate.SendChannel)

	var newTaskList []types.TaskInfo
	for _, param := range messageParamList {

		curTask := types.TaskInfo{
			MessageTemplateId: messageTemplate.ID,
			BusinessId:        taskHelper.GenerateBusinessId(messageTemplate.ID, messageTemplate.TemplateType),
			Receiver:          transformHelper.ArrayStringUniq(strings.Split(param.Receiver, ",")),
			IdType:            messageTemplate.IDType,
			SendChannel:       messageTemplate.SendChannel,
			TemplateType:      messageTemplate.TemplateType,
			MsgType:           messageTemplate.MsgType,
			ShieldType:        messageTemplate.ShieldType,
			ContentModel:      contentModel.BuilderContent(messageTemplate.Temp2Domain(), param),
			SendAccount:       messageTemplate.SendAccount,
			TemplateSn:        messageTemplate.TemplateSn,
			MessageParam:      param,
		}

		newTaskList = append(newTaskList, curTask)
	}
	sendTaskModel.TaskInfo = newTaskList
	return nil
}
