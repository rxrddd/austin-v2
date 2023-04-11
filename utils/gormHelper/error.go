package gormHelper

import (
	"errors"
	"gorm.io/gorm"
)

func IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
