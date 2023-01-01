package idType

const (
	UserId           int = 10
	Did              int = 20
	Phone            int = 30
	OpenId           int = 40
	Email            int = 50
	EnterpriseUserId int = 60
	DingDingUserId   int = 70
)

var (
	TypeDescription = map[int]string{
		UserId:           "userId",
		Did:              "did",
		Phone:            "phone",
		OpenId:           "openId",
		Email:            "email",
		EnterpriseUserId: "enterprise_user_id",
		DingDingUserId:   "ding_ding_user_id",
	}
)
