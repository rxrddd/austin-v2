package model

import "time"

type MsgRecord struct {
	ID                int64     `gorm:"column:id" json:"id"`
	MessageTemplateID int64     `gorm:"column:message_template_id" json:"message_template_id"` //消息模板ID
	RequestID         string    `gorm:"column:request_id" json:"request_id"`                   //唯一请求 ID
	Receiver          string    `gorm:"column:receiver" json:"receiver"`                       //接收人
	MsgId             string    `gorm:"column:msg_id" json:"msg_id"`                           //公众号消息id
	Channel           string    `gorm:"column:channel" json:"channel"`                         //渠道
	Msg               string    `gorm:"column:msg" json:"msg"`                                 //推送结果信息
	SendAt            string    `gorm:"column:send_at" json:"send_at"`                         //消息http 发送时间
	CreateAt          time.Time `gorm:"column:create_at" json:"create_at"`
	StartConsumeAt    string    `gorm:"column:start_consume_at" json:"start_consume_at"`     //开始消费时间
	EndConsumeAt      string    `gorm:"column:end_consume_at" json:"end_consume_at"`         //结束消费时间
	ConsumeSinceTime  string    `gorm:"column:consume_since_time" json:"consume_since_time"` //消费间距时间
	SendSinceTime     string    `gorm:"column:send_since_time" json:"send_since_time"`       //http->mq消费结束间距时间
	TaskInfo          string    `gorm:"column:task_info" json:"task_info"`
}

func (m MsgRecord) TableName() string {
	return "msg_record"
}
