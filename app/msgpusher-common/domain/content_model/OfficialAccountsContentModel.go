package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/jsonHelper"
)

type MiniProgram struct {
	Appid    string `json:"appid"`    //所需跳转到的小程序appid（该小程序 appid 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}
type OfficialAccountsContentModel struct {
	Data        map[string]string `json:"data"`        //消息数据
	Url         string            `json:"url"`         // 模板跳转链接
	TemplateSn  string            `json:"template_sn"` // 发送消息的模版ID
	MiniProgram MiniProgram       `json:"mini_program"`
}

func NewOfficialAccountsContentModel() *OfficialAccountsContentModel {
	return &OfficialAccountsContentModel{}
}

func (d OfficialAccountsContentModel) BuilderContent(messageTemplate *domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content OfficialAccountsContentModel
	jsonHelper.AnyToPtr(variables, &content)
	content.TemplateSn = messageTemplate.TemplateSn
	return content
}
