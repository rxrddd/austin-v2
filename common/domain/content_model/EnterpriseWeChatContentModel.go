package content_model

import (
	"austin-v2/common/domain"
	"austin-v2/pkg/types"
	"austin-v2/utils/taskHelper"
	"encoding/json"
)

type EnterpriseWeChatContentModel struct {
	SendType string `json:"sendType"`
	Content  string `json:"content"`
	MediaId  string `json:"mediaId"`
}

func NewEnterpriseWeChatContentModel() *EnterpriseWeChatContentModel {
	return &EnterpriseWeChatContentModel{}
}

func (d EnterpriseWeChatContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content EnterpriseWeChatContentModel
	_ = json.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.Content = taskHelper.ReplaceByMap(content.Content, newVariables)
	content.SendType = newVariables["sendType"]
	content.MediaId = newVariables["mediaId"]
	return content
}
