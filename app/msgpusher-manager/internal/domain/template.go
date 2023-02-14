package domain

type TemplateOneRequest struct {
	Id int64
}

type TemplateOneResp struct {
	ID                  int64
	Name                string
	AuditStatus         int
	IDType              int
	SendChannel         int
	TemplateType        int
	TemplateSn          string
	MsgType             int
	ShieldType          int
	MsgContent          string
	SendAccount         int64
	Creator             string
	Updator             string
	Auditor             string
	Team                string
	Proposer            string
	SmsChannel          string
	IsDeleted           int
	Created             int64
	Updated             int64
	DeduplicationConfig string
}

type TemplateListRequest struct {
	Name        string
	SendChannel string
	Page        int64
	PageSize    int64
}

type TemplateListRow struct {
	ID                  int64
	Name                string
	IdType              int64
	SendChannel         int64
	TemplateType        int64
	MsgType             int64
	ShieldType          int64
	MsgContent          string
	SendAccount         int64
	SendAccountName     string
	TemplateSn          string
	SmsChannel          string
	CreateAt            string
	DeduplicationConfig string
}

type TemplateListResp struct {
	Rows  []*TemplateListRow
	Total int64
}

type TemplateEditRequest struct {
	ID                  int64
	Name                string
	AuditStatus         int
	IDType              int
	SendChannel         int
	TemplateType        int
	TemplateSn          string
	MsgType             int
	ShieldType          int
	MsgContent          string
	SendAccount         int64
	Creator             string
	Updator             string
	Auditor             string
	Team                string
	Proposer            string
	SmsChannel          string
	IsDeleted           int
	Created             int64
	Updated             int64
	DeduplicationConfig string
}
