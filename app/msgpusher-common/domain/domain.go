package domain

import (
	"austin-v2/pkg/types"
)

type BuilderContent interface {
	BuilderContent(messageTemplate *MessageTemplate, messageParam types.MessageParam) interface{}
}

type MessageTemplate struct {
	ID                  int64  `json:"id"`                   //
	Name                string `json:"name"`                 // 标题
	AuditStatus         int    `json:"audit_status"`         // 当前消息审核状态： 10.待审核 20.审核成功 30.被拒绝
	IDType              int    `json:"id_type"`              // 消息的发送ID类型：10. userId 20.did 30.手机号 40.openId 50.email 60.企业微信userId
	SendChannel         int    `json:"send_channel"`         // 消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序 70.企业微信
	TemplateType        int    `json:"template_type"`        // 10.运营类 20.技术类接口调用
	TemplateSn          string `json:"template_sn"`          // 发送消息的模版ID
	MsgType             int    `json:"msg_type"`             // 10.通知类消息 20.营销类消息 30.验证码类消息
	ShieldType          int    `json:"shield_type"`          // 10.夜间不屏蔽 20.夜间屏蔽 30.夜间屏蔽(次日早上9点发送)
	MsgContent          string `json:"msg_content"`          // 消息内容 占位符用{$var}表示
	SendAccount         int64  `json:"send_account"`         // 发送账号 一个渠道下可存在多个账号
	Creator             string `json:"creator"`              // 创建者
	Updator             string `json:"updator"`              // 更新者
	Auditor             string `json:"auditor"`              // 审核人
	Team                string `json:"team"`                 // 业务方团队
	Proposer            string `json:"proposer"`             // 业务方
	SmsChannel          string `json:"sms_channel"`          // 短信渠道 send_channel=30的时候有用
	IsDeleted           int    `json:"is_deleted"`           // 是否删除：0.不删除 1.删除
	Created             int32  `json:"created"`              // 创建时间
	Updated             int32  `json:"updated"`              // 更新时间
	DeduplicationConfig string `json:"deduplication_config"` // 限流配置
}
