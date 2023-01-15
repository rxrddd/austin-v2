package mongo_model

type MsgRecord struct {
	ID                int64  `bson:"id" json:"id"`                                   //
	MessageTemplateID int64  `bson:"message_template_id" json:"message_template_id"` // 消息模板ID
	RequestID         string `bson:"request_id" json:"request_id"`                   // 唯一请求 ID
	CreateAt          string `bson:"create_at" json:"create_at"`
	TaskInfo          string `bson:"task_info" json:"task_info"`
	Receiver          string `bson:"receiver" json:"receiver"`
	MsgId             string `bson:"msg_id" json:"msg_id"`
	Channel           string `bson:"channel" json:"channel"`
	Msg               string `bson:"msg" json:"msg"`
	StartConsumeAt    string `bson:"start_consume_at" json:"start_consume_at"`
	EndConsumeAt      string `bson:"end_consume_at" json:"end_consume_at"`
	SendAt            string `bson:"send_at" json:"send_at"`
	ConsumeSinceTime  string `bson:"consume_since_time" json:"consume_since_time"`
	SendSinceTime     string `bson:"send_since_time" json:"send_since_time"`
}
