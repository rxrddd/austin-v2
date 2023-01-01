package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
)

type ImContentModel struct {
}

func NewImContentModel() *ImContentModel {
	return &ImContentModel{}
}

func (m ImContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	return m
}
