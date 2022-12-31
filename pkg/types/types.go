package types

import "context"

type TaskInfo struct {
	MessageTemplateId int64        `json:"messageTemplateId"`
	BusinessId        int64        `json:"businessId"`
	Receiver          []string     `json:"receiver"` //先去重
	IdType            int          `json:"idType"`
	SendChannel       int          `json:"sendChannel"`
	TemplateType      int          `json:"templateType"`
	MsgType           int          `json:"msgType"`
	ShieldType        int          `json:"shieldType"`
	ContentModel      interface{}  `json:"contentModel"`
	SendAccount       int          `json:"sendAccount"`                           //发消息使用的账号
	TemplateSn        string       `json:"templateSn"`                            // 发送消息的模版ID
	SmsChannel        string       `gorm:"column:sms_channel" json:"sms_channel"` // 短信渠道 send_channel=30的时候有用
	MessageParam      MessageParam `json:"messageParamList"`
}

type ContentModel struct {
}

type SendTaskModel struct {
	MessageTemplateId int64          `json:"messageTemplateId"`
	MessageParamList  []MessageParam `json:"messageParamList"`
	TaskInfo          []TaskInfo     `json:"taskInfo"`
}

type MessageParam struct {
	Receiver  string                 `json:"receiver"`           //接收者 多个用,逗号号分隔开
	Variables map[string]interface{} `json:"variables,optional"` //可选 消息内容中的可变部分(占位符替换)
	Extra     map[string]interface{} `json:"extra,optional"`     //可选 扩展参数
}

type IHandler interface {
	Name() string
	Execute(ctx context.Context, taskInfo *TaskInfo) error
	Allow(ctx context.Context) bool
}
