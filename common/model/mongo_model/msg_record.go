package mongo_model

import "time"

type MsgRecord struct {
	ID                int64     `bson:"id" json:"id"`                                   //
	MessageTemplateID int64     `bson:"message_template_id" json:"message_template_id"` // 消息模板ID
	RequestID         string    `bson:"request_id" json:"request_id"`                   // 唯一请求 ID
	Receiver          string    `bson:"receiver" json:"receiver"`                       //接收人
	MsgId             string    `bson:"msg_id" json:"msg_id"`                           //公众号消息id
	Channel           string    `bson:"channel" json:"channel"`                         //渠道
	Msg               string    `bson:"msg" json:"msg"`                                 //推送结果信息
	SendAt            string    `bson:"send_at" json:"send_at"`                         //消息http 发送时间
	CreateAt          time.Time `bson:"create_at" json:"create_at"`
	StartConsumeAt    string    `bson:"start_consume_at" json:"start_consume_at"`     //开始消费时间
	EndConsumeAt      string    `bson:"end_consume_at" json:"end_consume_at"`         //结束消费时间
	ConsumeSinceTime  string    `bson:"consume_since_time" json:"consume_since_time"` //消费间距时间
	SendSinceTime     string    `bson:"send_since_time" json:"send_since_time"`       //http->mq消费结束间距时间
	TaskInfo          string    `bson:"task_info" json:"task_info"`
}
