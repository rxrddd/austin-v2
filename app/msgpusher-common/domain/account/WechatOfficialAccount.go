package account

type WechatOfficialAccount struct {
	OpenId        string            `json:"openId"`
	TemplateId    string            `json:"templateId"`
	Url           string            `json:"url"`
	MiniProgramId string            `json:"miniProgramId"`
	Path          string            `json:"path"`
	Map           map[string]string `json:"map"`
}
