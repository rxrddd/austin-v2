package process

import (
	"austin-v2/app/msgpusher/internal/data/model"
	"austin-v2/common/domain/content_model"
	"austin-v2/pkg/types"
	"austin-v2/utils/taskHelper"
	"austin-v2/utils/transformHelper"
	"context"
	"strings"
	"time"
)

type AssembleAction struct {
	//uc *biz.MessageTemplateUseCase
}

func NewAssembleAction(
//uc *biz.MessageTemplateUseCase,
) *AssembleAction {
	return &AssembleAction{
		//uc: uc,
	}
}

func (p *AssembleAction) Process(ctx context.Context, sendTaskModel *types.SendTaskModel, messageTemplate model.MessageTemplate) error {
	messageParamList := sendTaskModel.MessageParamList
	contentModel := content_model.GetBuilderContentBySendChannel(messageTemplate.SendChannel)

	var newTaskList []types.TaskInfo
	for _, param := range messageParamList {
		curTask := types.TaskInfo{
			RequestId:         sendTaskModel.RequestId,
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
			SendAt:            time.Now(),
			MessageParam:      param,
		}

		newTaskList = append(newTaskList, curTask)
	}
	sendTaskModel.TaskInfo = newTaskList
	return nil
}
