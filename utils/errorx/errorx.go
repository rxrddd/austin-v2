package errorx

import "github.com/go-kratos/kratos/v2/errors"

const BusinessError = "BUSINESS_ERROR"

func New(code int, reason, message string) *errors.Error {
	return errors.New(code, reason, message)
}
func NewBizErr(message string) *errors.Error {
	return errors.New(0, BusinessError, message)
}
func NewBizReason(reason, message string) *errors.Error {
	return errors.New(0, reason, message)
}
