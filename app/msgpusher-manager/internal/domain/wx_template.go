package domain

type OfficialAccountTemplateRequest struct {
	SendAccount int64 `json:"send_account"`
}

type OfficialAccountTemplateRow struct {
	TemplateID string `json:"template_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Example    string `json:"example"`
}
type OfficialAccountTemplateResp struct {
	Rows []*OfficialAccountTemplateRow
}
