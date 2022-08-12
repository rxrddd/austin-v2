package util

import (
	"github.com/ZQCard/kratos-base-project/pkg/utils/encryption"
	"github.com/ZQCard/kratos-base-project/pkg/utils/stringHelper"
)

// HashPassword 密码加密
func HashPassword(password string) (salt string, dbPassword string) {
	salt = encryption.EncodeMD5(stringHelper.RandString(10))
	dbPassword = encryption.EncodeMD5(password + salt)
	return
}

// HashSaltPassword 密码加密
func HashSaltPassword(salt string, password string) (dbPassword string) {
	dbPassword = encryption.EncodeMD5(password + salt)
	return
}

// CheckPasswordHash 验证密码
func CheckPasswordHash(password, salt, originPassword string) bool {
	return originPassword == encryption.EncodeMD5(password+salt)
}
