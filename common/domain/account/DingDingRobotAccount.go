package account

type DingDingRobotAccount struct {
	AccessToken string `json:"access_token"` //密钥
	Secret      string `json:"secret"`       //密钥
	//Webhook string `json:"webhook"` //自定义群机器人中的 webhook
}
