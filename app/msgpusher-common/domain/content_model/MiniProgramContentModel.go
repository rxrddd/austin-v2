package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
)

type MiniProgramContentModel struct {
}

func NewMiniProgramContentModel() *MiniProgramContentModel {
	return &MiniProgramContentModel{}
}

func (m MiniProgramContentModel) BuilderContent(messageTemplate domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	return m
}
