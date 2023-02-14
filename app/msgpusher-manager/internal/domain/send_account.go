package domain

type SendAccountListRequest struct {
	Title       string
	SendChannel string
	Page        int64
	PageSize    int64
}
type SendAccountRow struct {
	ID          int64
	Title       string
	Config      string
	SendChannel string
	Status      int
}
type SendAccountListResp struct {
	Rows  []*SendAccountRow
	Total int64
}
type SendAccountQueryResp struct {
	Rows []*SendAccountRow
}

type SendAccountEditRequest struct {
	ID          int64
	SendChannel string
	Config      string
	Title       string
}
