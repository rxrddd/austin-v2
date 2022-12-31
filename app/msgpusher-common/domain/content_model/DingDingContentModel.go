package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"encoding/json"
)

type DingDingContentModel struct {
	//SendType string `json:"sendType"`
	Content string `json:"content"`
	//MediaId  string `json:"mediaId"`
}

func NewDingDingContentModel() *DingDingContentModel {
	return &DingDingContentModel{}
}

func (d DingDingContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content DingDingContentModel
	_ = json.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.Content = taskHelper.ReplaceByMap(content.Content, newVariables)
	//content.SendType = newVariables["sendType"]
	//content.MediaId = newVariables["mediaId"]
	return content
}
