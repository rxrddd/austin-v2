package errResponse

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

var reasonMessageAll = map[string]string{
	"SUCCESS":       "请求成功",
	"UNKNOWN_ERROR": "未知错误",

	"MISSING_PARAMS":       "缺少搜索参数",
	"MISSING_CONDITION_ID": "id不得为空",

	"ADMINISTRATOR_RECORD_NOT_FOUND": "管理员数据不存在",
	"ADMINISTRATOR_PASSWORD_ERROR":   "管理员密码错误",
	"ADMINISTRATOR_FORBIDDEN":        "管理员已禁用",
	"ADMINISTRATOR_DELETED":          "管理员已删除",
	"ADMINISTRATOR_USERNAME_EXIST":   "管理员用户名已存在",
	"ADMINISTRATOR_MOBILE_EXIST":     "管理员手机号已存在",

	"UNAUTHORIZED": "账号未登陆",

	"GOODS_RECORD_NOT_FOUND": "商品数据不存在",
	"GOODS_RECORD_DELETED":   "商品已删除",
	"GOODS_NOT_ENOUGH_STOCK": "商品库存不足",

	"SYSTEM_ERROR":            "系统繁忙,请稍后再试",
	"SERVICE_GATEWAY_TIMEOUT": "服务不可达",
}

var reasonCodeAll = map[string]int{
	"SUCCESS":       0,
	"UNKNOWN_ERROR": 1,

	"MISSING_CONDITION":    10000,
	"MISSING_CONDITION_ID": 10001,

	"ADMINISTRATOR_RECORD_NOT_FOUND": 20001,
	"ADMINISTRATOR_PASSWORD_ERROR":   20002,
	"ADMINISTRATOR_FORBIDDEN":        20003,
	"ADMINISTRATOR_DELETED":          20004,
	"ADMINISTRATOR_USERNAME_EXIST":   20005,
	"ADMINISTRATOR_MOBILE_EXIST":     20006,

	"GOODS_RECORD_NOT_FOUND": 30001,
	"GOODS_RECORD_DELETED":   30002,
	"GOODS_NOT_ENOUGH_STOCK": 30003,

	"UNAUTHORIZED": 4000,

	"SYSTEM_ERROR":            50000,
	"SERVICE_GATEWAY_TIMEOUT": 50001,
}

// SetCustomizeErrInfo 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfo(err error) error {
	fmt.Println("err")
	fmt.Println(err)
	e := errors.FromError(err)
	// 如果 e.Code = 504 则是服务不可达
	if e.Code == http.StatusGatewayTimeout {
		return SetCustomizeErrInfoByReason("SERVICE_GATEWAY_TIMEOUT")
	}
	reason := e.Reason
	if reason == "" {
		reason = "UNKNOWN_ERROR"
	}
	if _, ok := reasonCodeAll[reason]; !ok {
		return err
	}
	return SetCustomizeErrInfoByReason(e.Reason)
}

// SetCustomizeErrInfoByReason 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfoByReason(reason string) error {
	code, message := reasonCodeAll[reason], reasonMessageAll[reason]
	return errors.New(code, reason, message)
}
