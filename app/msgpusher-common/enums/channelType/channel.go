package channelType

const (
	//Im                 int = 10
	//Push               int = 20
	Sms                int = 30 //短信
	Email              int = 40 //邮箱
	OfficialAccounts   int = 50 //公众号
	MiniProgram        int = 60 //小程序
	EnterpriseWeChat   int = 70 //企业微信
	DingDingRobot      int = 80 //钉钉机器人
	DingDingWorkNotice int = 90 //钉钉工作通知
)

var (
	TypeText = map[int]string{
		//Im:                 "IM(站内信)",
		//Push:               "push(通知栏)",
		Sms:                "短信",
		Email:              "邮件",
		OfficialAccounts:   "服务号",
		MiniProgram:        "小程序",
		EnterpriseWeChat:   "企业微信",
		DingDingRobot:      "钉钉机器人",
		DingDingWorkNotice: "钉钉工作通知",
	}
	TypeCodeEn = map[int]string{
		//Im:                 "im",
		//Push:               "push",
		Sms:                "sms",
		Email:              "email",
		OfficialAccounts:   "official_accounts",
		MiniProgram:        "mini_program",
		EnterpriseWeChat:   "enterprise_we_chat",
		DingDingRobot:      "ding_ding_robot",
		DingDingWorkNotice: "ding_ding_work_notice",
	}
	TypeEnCode = map[string]int{
		//"im":                    Im,
		//"push":                  Push,
		"sms":                   Sms,
		"email":                 Email,
		"official_accounts":     OfficialAccounts,
		"mini_program":          MiniProgram,
		"enterprise_we_chat":    EnterpriseWeChat,
		"ding_ding_robot":       DingDingRobot,
		"ding_ding_work_notice": DingDingWorkNotice,
	}
)
