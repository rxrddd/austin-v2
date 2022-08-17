package errResponse

import (
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

const ReasonSuccess = "SUCCESS"
const ReasonUnknownError = "UNKNOWN_ERROR"

const ReasonMissingParams = "MISSING_PARAMS"
const ReasonMissingId = "MISSING_ID"
const ReasonParamsError = "PARAMS_ERROR"

const ReasonAdministratorNotFound = "ADMINISTRATOR_NOT_FOUND"
const ReasonAdministratorPasswordError = "ADMINISTRATOR_PASSWORD_ERROR"
const ReasonAdministratorForbidden = "ADMINISTRATOR_FORBIDDEN"
const ReasonAdministratorDeleted = "ADMINISTRATOR_DELETED"
const ReasonAdministratorUsernameExist = "ADMINISTRATOR_USERNAME_EXIST"
const ReasonAdministratorMobileExist = "ADMINISTRATOR_MOBILE_EXIST"

const ReasonAdministratorUNAUTHORIZED = "UNAUTHORIZED"

const ReasonSystemError = "SYSTEM_ERROR"
const ReasonServiceGatewayTimeout = "SERVICE_GATEWAY_TIMEOUT"

var reasonMessageAll = map[string]string{
	ReasonSuccess:      "请求成功",
	ReasonUnknownError: "未知错误",

	ReasonParamsError:   "请求参数错误",
	ReasonMissingParams: "缺少搜索参数",
	ReasonMissingId:     "id不得为空",

	ReasonAdministratorNotFound:      "管理员数据不存在",
	ReasonAdministratorPasswordError: "管理员密码错误",
	ReasonAdministratorForbidden:     "管理员已禁用",
	ReasonAdministratorDeleted:       "管理员已删除",
	ReasonAdministratorUsernameExist: "管理员用户名已存在",
	ReasonAdministratorMobileExist:   "管理员手机号已存在",

	ReasonAdministratorUNAUTHORIZED: "管理员未登陆",

	ReasonSystemError:           "系统繁忙,请稍后再试",
	ReasonServiceGatewayTimeout: "服务不可达",
}

var reasonCodeAll = map[string]int{
	ReasonSuccess:      0,
	ReasonUnknownError: 1,

	ReasonParamsError:   10000,
	ReasonMissingParams: 10001,
	ReasonMissingId:     10002,

	ReasonAdministratorNotFound:      20001,
	ReasonAdministratorPasswordError: 20002,
	ReasonAdministratorForbidden:     20003,
	ReasonAdministratorDeleted:       20004,
	ReasonAdministratorUsernameExist: 20005,
	ReasonAdministratorMobileExist:   20006,

	ReasonAdministratorUNAUTHORIZED: 4000,

	ReasonSystemError:           50000,
	ReasonServiceGatewayTimeout: 50001,
}

var reasonGrpcCodeAll = map[string]int{
	ReasonSuccess:      http.StatusOK,
	ReasonUnknownError: http.StatusBadRequest,

	ReasonParamsError:   http.StatusBadRequest,
	ReasonMissingParams: http.StatusBadRequest,
	ReasonMissingId:     http.StatusBadRequest,

	ReasonAdministratorNotFound:      http.StatusBadRequest,
	ReasonAdministratorPasswordError: http.StatusBadRequest,
	ReasonAdministratorForbidden:     http.StatusBadRequest,
	ReasonAdministratorDeleted:       http.StatusBadRequest,
	ReasonAdministratorUsernameExist: http.StatusBadRequest,
	ReasonAdministratorMobileExist:   http.StatusBadRequest,

	ReasonAdministratorUNAUTHORIZED: http.StatusUnauthorized,

	ReasonSystemError:           http.StatusInternalServerError,
	ReasonServiceGatewayTimeout: http.StatusGatewayTimeout,
}

// SetCustomizeErrInfo 根据err.Reason返回自定义包装错误
func SetCustomizeErrInfo(err error) error {
	e := errors.FromError(err)
	// 如果 e.Code = 504 则是服务不可达
	if e.Code == http.StatusGatewayTimeout {
		return SetCustomizeErrInfoByReason(ReasonServiceGatewayTimeout)
	}
	reason := e.Reason
	if reason == "" {
		reason = ReasonUnknownError
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

// GetErrInfoByReason 根据err.Reason返回自定义包装错误
func GetErrInfoByReason(reason string) string {
	return reasonMessageAll[reason]
}

// SetGRpcErrByReason 根据err.Reason返回自定义包装错误
func SetGRpcErrByReason(reason string) error {
	code, message := reasonGrpcCodeAll[reason], reasonMessageAll[reason]
	return errors.New(code, reason, message)
}
