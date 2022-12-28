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
const TimeFormatError = "TIME_FORMAT_ERROR"
const ReasonRecordNotFound = "RECORD_NOT_FOUND"

const ReasonAdministratorNotFound = "ADMINISTRATOR_NOT_FOUND"
const ReasonAdministratorPasswordError = "ADMINISTRATOR_PASSWORD_ERROR"
const ReasonAdministratorForbidden = "ADMINISTRATOR_FORBIDDEN"
const ReasonAdministratorDeleted = "ADMINISTRATOR_DELETED"
const ReasonAdministratorUsernameExist = "ADMINISTRATOR_USERNAME_EXIST"
const ReasonAdministratorMobileExist = "ADMINISTRATOR_MOBILE_EXIST"

const ReasonUnauthorizedRole = "UNAUTHORIZED_ROLE"
const ReasonUnauthorizedInfoMissing = "UNAUTHORIZED_INFO_MISSING"
const ReasonAuthorizedDataMissing = "AUTHORIZATION_DATA_MISSING"
const ReasonAuthorizationRoleExist = "AUTHORIZATION_ROLE_EXIST"
const ReasonAuthorizationRoleNotFound = "AUTHORIZATION_ROLE_NOT_FOUND"
const ReasonAuthorizationUserHasRoleAlready = "AUTHORIZATION_USE_HAS_ROLE_ALREADY"

const ReasonAuthorizationApiNotFound = "AUTHORIZATION_API_NOT_FOUND"
const ReasonAuthorizationApiExist = "AUTHORIZATION_API_EXIST"

const ReasonFileNameMissing = "FILE_NAME_MISSING"
const ReasonFileMissing = "FILE_MISSING"
const ReasonFileOverLimitSize = "FILE_OVER_LIMIT_SIZE"
const ReasonOssConfigWrong = "OSS_CONFIG_WRONG"
const ReasonOssBucketWrong = "OSS_BUCKET_WRONG"
const ReasonOssPutObjectFail = "OSS_PUT_OBJECT_FILE"

const ReasonAdministratorUnauthorized = "UNAUTHORIZED"

const ReasonSystemError = "SYSTEM_ERROR"
const ReasonServiceGatewayTimeout = "SERVICE_GATEWAY_TIMEOUT"
const ReasonServiceUnavailable = "SERVICE_GATEWAY_UNAVAILABLE"

var reasonMessageAll = map[string]string{
	ReasonSuccess:      "请求成功",
	ReasonUnknownError: "未知错误",

	ReasonParamsError:    "请求参数错误",
	ReasonMissingParams:  "请求参数错误",
	TimeFormatError:      "时间格式错误",
	ReasonMissingId:      "id不得为空",
	ReasonRecordNotFound: "数据不存在",

	ReasonAdministratorNotFound:      "管理员数据不存在",
	ReasonAdministratorPasswordError: "管理员密码错误",
	ReasonAdministratorForbidden:     "管理员已禁用",
	ReasonAdministratorDeleted:       "管理员已删除",
	ReasonAdministratorUsernameExist: "管理员用户名已存在",
	ReasonAdministratorMobileExist:   "管理员手机号已存在",

	ReasonAdministratorUnauthorized: "管理员未登陆",
	ReasonUnauthorizedInfoMissing:   "角色授权信息不存在",
	ReasonAuthorizedDataMissing:     "权限数据不存在",
	ReasonUnauthorizedRole:          "角色未授权",

	ReasonAuthorizationRoleExist:          "角色已存在",
	ReasonAuthorizationRoleNotFound:       "角色不存在",
	ReasonAuthorizationUserHasRoleAlready: "用户已经拥有角色",

	ReasonAuthorizationApiNotFound: "API不存在",
	ReasonAuthorizationApiExist:    "API已存在",

	ReasonFileNameMissing:   "文件名称不得为空",
	ReasonFileMissing:       "文件不得为空",
	ReasonFileOverLimitSize: "文件超过最大限制",

	ReasonOssConfigWrong:   "OSS配置错误",
	ReasonOssBucketWrong:   "OSS Bucket配置错误",
	ReasonOssPutObjectFail: "OSS文件上传失败",

	ReasonSystemError:           "系统繁忙,请稍后再试",
	ReasonServiceGatewayTimeout: "服务不可达",
	ReasonServiceUnavailable:    "服务不可达",
}

var reasonCodeAll = map[string]int{
	ReasonSuccess:      0,
	ReasonUnknownError: 1,

	ReasonParamsError:    10000,
	ReasonMissingParams:  10001,
	ReasonMissingId:      10002,
	ReasonRecordNotFound: 10003,
	TimeFormatError:      10004,

	ReasonAdministratorNotFound:      20001,
	ReasonAdministratorPasswordError: 20002,
	ReasonAdministratorForbidden:     20003,
	ReasonAdministratorDeleted:       20004,
	ReasonAdministratorUsernameExist: 20005,
	ReasonAdministratorMobileExist:   20006,

	ReasonAdministratorUnauthorized:       40000,
	ReasonUnauthorizedRole:                40001,
	ReasonUnauthorizedInfoMissing:         40002,
	ReasonAuthorizedDataMissing:           40003,
	ReasonAuthorizationRoleExist:          40200,
	ReasonAuthorizationRoleNotFound:       40204,
	ReasonAuthorizationUserHasRoleAlready: 40202,

	ReasonAuthorizationApiExist:    40302,
	ReasonAuthorizationApiNotFound: 40304,

	ReasonFileNameMissing:   40401,
	ReasonFileMissing:       40402,
	ReasonFileOverLimitSize: 40403,
	ReasonOssConfigWrong:    40404,
	ReasonOssBucketWrong:    40405,
	ReasonOssPutObjectFail:  40406,

	ReasonSystemError:           50000,
	ReasonServiceGatewayTimeout: 50004,
	ReasonServiceUnavailable:    50003,
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

	ReasonAdministratorUnauthorized: http.StatusUnauthorized,

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
	// 如果 e.Code = 503 则是服务连接被拒绝
	if e.Code == http.StatusServiceUnavailable {
		return SetCustomizeErrInfoByReason(ReasonServiceUnavailable)
	}
	reason := e.Reason
	if reason == "" {
		reason = ReasonUnknownError
	}
	if _, ok := reasonCodeAll[reason]; !ok {
		return err
	}
	// 如果是参数错误， 则检查err是否有值， 有则直接参会
	if e.Reason == ReasonParamsError && e.Message != "" {
		return errors.New(reasonCodeAll[reason], reason, e.Message)
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

// SetErrByReason 根据err.Reason返回自定义包装错误
func SetErrByReason(reason string) error {
	code, message := reasonGrpcCodeAll[reason], reasonMessageAll[reason]
	return errors.New(code, reason, message)
}
