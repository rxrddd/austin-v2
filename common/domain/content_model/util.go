package content_model

import (
	"austin-v2/common/domain"
	"austin-v2/common/enums/channelType"
)

var contentMap = map[int]domain.BuilderContent{
	//channelType.Im:                 NewImContentModel(),               //IM(站内信)
	//channelType.Push:               NewPushContentModel(),             //push(通知栏)
	channelType.Sms:                NewSmsContentModel(),              //sms(短信)
	channelType.Email:              NewEmailContentModel(),            //email(邮件)
	channelType.OfficialAccounts:   NewOfficialAccountsContentModel(), //OfficialAccounts(服务号)
	channelType.MiniProgram:        NewMiniProgramContentModel(),      //miniProgram(小程序)
	channelType.EnterpriseWeChat:   NewEnterpriseWeChatContentModel(), //EnterpriseWeChat(企业微信)
	channelType.DingDingRobot:      NewDingDingContentModel(),         //dingDingRobot(钉钉机器人)
	channelType.DingDingWorkNotice: NewDingDingContentModel(),         //dingDingWorkNotice(钉钉工作通知)
}

// GetBuilderContentBySendChannel 消息发送渠道：10.IM 20.Push 30.短信 40.Email 50.公众号 60.小程序 70.企业微信
func GetBuilderContentBySendChannel(sendChannel int) domain.BuilderContent {
	return contentMap[sendChannel]
}
