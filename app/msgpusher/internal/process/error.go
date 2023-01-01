package process

import "errors"

var (
	sendErr         = errors.New("发送消息错误")
	clientParamsErr = errors.New("客户端参数错误")
)
