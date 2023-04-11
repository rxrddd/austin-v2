// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMsgRecord = "msg_record"

// MsgRecord mapped from table <msg_record>
type MsgRecord struct {
	ID                int64     `gorm:"column:id;primaryKey" json:"id"`
	MessageTemplateID int64     `gorm:"column:message_template_id;not null" json:"message_template_id"` // 消息模板ID
	RequestID         string    `gorm:"column:request_id;not null" json:"request_id"`                   // 唯一请求 ID
	Receiver          string    `gorm:"column:receiver;not null" json:"receiver"`                       // 接收人
	MsgID             string    `gorm:"column:msg_id;not null" json:"msg_id"`                           // 公众号消息id
	Channel           string    `gorm:"column:channel;not null" json:"channel"`                         // 渠道
	Msg               string    `gorm:"column:msg;not null" json:"msg"`                                 // 推送结果信息
	SendAt            string    `gorm:"column:send_at;not null" json:"send_at"`                         // 消息http 发送时间
	CreateAt          time.Time `gorm:"column:create_at;not null" json:"create_at"`
	StartConsumeAt    string    `gorm:"column:start_consume_at;not null" json:"start_consume_at"`     // 开始消费时间
	EndConsumeAt      string    `gorm:"column:end_consume_at;not null" json:"end_consume_at"`         // 结束消费时间
	ConsumeSinceTime  string    `gorm:"column:consume_since_time;not null" json:"consume_since_time"` // 消费间距时间
	SendSinceTime     string    `gorm:"column:send_since_time;not null" json:"send_since_time"`       // http->mq消费结束间距时间
	TaskInfo          string    `gorm:"column:task_info;not null" json:"task_info"`
}

// TableName MsgRecord's table name
func (*MsgRecord) TableName() string {
	return TableNameMsgRecord
}