package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
)

type PushContentModel struct {
}

func NewPushContentModel() *PushContentModel {
	return &PushContentModel{}
}

func (d PushContentModel) BuilderContent(messageTemplate domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	return d
}
