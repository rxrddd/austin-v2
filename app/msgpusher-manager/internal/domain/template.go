package domain

type TemplateOneRequest struct {
	Id int64
}

type TemplateOneResp struct {
	ID                  int64
	Name                string
	AuditStatus         int32
	IDType              int32
	SendChannel         int32
	TemplateType        int32
	TemplateSn          string
	MsgType             int32
	ShieldType          int32
	MsgContent          string
	SendAccount         int32
	CreateBy            string
	UpdateBy            string
	SmsChannel          string
	Status              int32
	CreateAt            int64
	UpdateAt            int64
	DeduplicationConfig string
}

type TemplateListRequest struct {
	Name        string
	SendChannel string
	PageNo      int64
	PageSize    int64
}

type TemplateListRow struct {
	ID                  int64
	Name                string
	IdType              int32
	SendChannel         int32
	TemplateType        int32
	MsgType             int32
	ShieldType          int32
	MsgContent          string
	SendAccount         int32
	SendAccountName     string
	TemplateSn          string
	SmsChannel          string
	CreateAt            string
	DeduplicationConfig string
}

type TemplateListResp struct {
	Rows  []*TemplateListRow
	Total int32
}

type TemplateEditRequest struct {
	ID                  int64
	Name                string
	AuditStatus         int32
	IDType              int32
	SendChannel         int32
	TemplateType        int32
	TemplateSn          string
	MsgType             int32
	ShieldType          int32
	MsgContent          string
	SendAccount         int32
	CreateBy            string
	UpdateBy            string
	SmsChannel          string
	Status              int32
	CreateAt            int64
	UpdateAt            int64
	DeduplicationConfig string
}
