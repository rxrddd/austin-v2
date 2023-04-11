package content_model

import (
	"austin-v2/common/domain"
	"austin-v2/pkg/types"
	"austin-v2/utils/taskHelper"
	"encoding/json"
)

type EmailContentModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewEmailContentModel() *EmailContentModel {
	return &EmailContentModel{}
}

func (d EmailContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content EmailContentModel
	_ = json.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.Content = taskHelper.ReplaceByMap(content.Content, newVariables)
	content.Title = newVariables["title"]
	return content
}
