package domain

type SmsRecordRequest struct {
	TemplateId  string
	RequestId   string
	SendChannel string
	Page        int64
	PageSize    int64
}
type MsgRecordRequest struct {
	TemplateId string
	RequestId  string
	Channel    string
	PageNo     int64
	PageSize   int64
}
type MsgRecordRow struct {
	MessageTemplateId int64
	RequestId         string
	Receiver          string
	MsgId             string
	Channel           string
	Msg               string
	SendAt            string
	CreateAt          string
	SendSinceTime     string
	ID                int64
}
type MsgRecordResp struct {
	Rows  []*MsgRecordRow
	Total int32
}
