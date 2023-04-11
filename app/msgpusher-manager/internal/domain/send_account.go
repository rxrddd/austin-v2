package domain

type SendAccountListRequest struct {
	Title       string
	SendChannel string
	PageNo      int64
	PageSize    int64
}
type SendAccountRow struct {
	ID          int32
	Title       string
	Config      string
	SendChannel string
	Status      int32
}
type SendAccountListResp struct {
	Rows  []*SendAccountRow
	Total int32
}
type SendAccountQueryResp struct {
	Rows []*SendAccountRow
}

type SendAccountEditRequest struct {
	ID          int32
	SendChannel string
	Config      string
	Title       string
}
