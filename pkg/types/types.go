package types

import (
	"context"
	"encoding/json"
)

type TaskInfo struct {
	RequestId         string       `json:"request_id"`
	MessageTemplateId int64        `json:"message_template_id"`
	BusinessId        int64        `json:"business_id"`
	Receiver          []string     `json:"receiver"` //先去重
	IdType            int          `json:"id_type"`
	SendChannel       int          `json:"send_channel"`
	TemplateType      int          `json:"template_type"`
	MsgType           int          `json:"msg_type"`
	ShieldType        int          `json:"shield_type"`
	ContentModel      interface{}  `json:"content_model"`
	SendAccount       int64        `json:"send_account"`
	TemplateSn        string       `json:"template_sn"`
	SmsChannel        string       `json:"sms_channel"`
	MessageParam      MessageParam `json:"message_param"`
}

func (t *TaskInfo) String() string {
	marshal, _ := json.Marshal(t)
	return string(marshal)
}

type ContentModel struct {
}

type SendTaskModel struct {
	RequestId         string         `json:"request_id"`
	MessageTemplateId int64          `json:"message_template_id"`
	MessageParamList  []MessageParam `json:"message_param_list"`
	TaskInfo          []TaskInfo     `json:"task_info"`
}

type MessageParam struct {
	Receiver  string                 `json:"receiver"`           //接收者 多个用,逗号号分隔开
	Variables map[string]interface{} `json:"variables,optional"` //可选 消息内容中的可变部分(占位符替换)
	Extra     map[string]interface{} `json:"extra,optional"`     //可选 扩展参数
}

type IHandler interface {
	Name() string
	Execute(ctx context.Context, taskInfo *TaskInfo) error
	Allow(ctx context.Context, taskInfo *TaskInfo) bool
}
type ISmsScript interface {
	Name() string
	Send(ctx context.Context, taskInfo *TaskInfo) error
}
