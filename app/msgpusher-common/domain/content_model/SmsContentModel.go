package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"encoding/json"
)

type SmsContentModel struct {
	Content        string `json:"content"`         //原始模板 您的验证码是{$code}，{$min}分钟内有效。请勿向他人泄露。如果非本人操作，可忽略本消息。
	ReplaceContent string `json:"replace_content"` //替换后的模板 您的验证码是1011，15分钟内有效。请勿向他人泄露。如果非本人操作，可忽略本消息。
	Url            string `json:"url"`
}

func NewSmsContentModel() *SmsContentModel {
	return &SmsContentModel{}
}

/**
messageParam 入参
messageTemplate 模板数据库配置
*/
func (s SmsContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content SmsContentModel
	_ = json.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	content.ReplaceContent = taskHelper.ReplaceByMap(content.Content, newVariables)
	if v, ok := newVariables["url"]; ok && v != "" {
		content.Url = taskHelper.GenerateUrl(v, messageTemplate.ID, messageTemplate.TemplateType)
	}
	return content
}
