package content_model

import (
	"austin-v2/app/msgpusher-common/domain"
	"austin-v2/pkg/types"
	"austin-v2/pkg/utils/taskHelper"
	"encoding/json"

	"github.com/spf13/cast"
)

type OfficialAccountsContentModel struct {
	Map        map[string]string `json:"map"`         //消息数据
	Url        string            `json:"url"`         // 消息的URL地址
	TemplateId string            `json:"template_sn"` // 发送消息的模版ID
}

func NewOfficialAccountsContentModel() *OfficialAccountsContentModel {
	return &OfficialAccountsContentModel{}
}

func (d OfficialAccountsContentModel) BuilderContent(messageTemplate domain.MessageTemplate, messageParam types.MessageParam) interface{} {
	variables := messageParam.Variables
	var content OfficialAccountsContentModel
	_ = json.Unmarshal([]byte(messageTemplate.MsgContent), &content)
	newVariables := getStringVariables(variables)
	if v, ok := newVariables["url"]; ok && v != "" {
		content.Url = taskHelper.GenerateUrl(v, messageTemplate.ID, messageTemplate.TemplateType)
	}
	content.Map = cast.ToStringMapString(variables["map"])
	content.TemplateId = messageTemplate.TemplateSn
	return content
}
