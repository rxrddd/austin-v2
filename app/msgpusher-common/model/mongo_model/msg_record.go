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
}
