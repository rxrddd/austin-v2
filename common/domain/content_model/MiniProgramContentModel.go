package content_model

import (
	"austin-v2/common/domain"
	"austin-v2/pkg/types"
	"austin-v2/utils/jsonHelper"
)

type MiniProgramContentModel struct {
	Data             map[string]string `json:"data"`              //消息数据
	TemplateSn       string            `json:"template_sn"`       // 发送消息的模版ID
	Page             string            `json:"page"`              // 发送消息的模版ID
	MiniProgramState string            `json:"miniprogram_state"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string            `json:"lang"`              // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

func NewMiniProgramContentModel() *MiniProgramContentModel {
	return &MiniProgramContentModel{}
}

func (d MiniProgramContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content MiniProgramContentModel
	jsonHelper.AnyToPtr(variables, &content)
	content.TemplateSn = messageTemplate.TemplateSn
	return content
}
