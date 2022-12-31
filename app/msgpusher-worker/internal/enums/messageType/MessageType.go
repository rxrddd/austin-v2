package messageType

const (
	Notice    = 10
	Marketing = 20
	AuthCode  = 30
)

var TypeCodeEn = map[int]string{
	Notice:    "notice",
	Marketing: "marketing",
	AuthCode:  "auth_code",
}
