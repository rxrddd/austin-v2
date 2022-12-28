package encryption

import (
	"austin-v2/pkg/utils/stringHelper"
)

// HashPassword 密码加密
func HashPassword(password string) (salt string, dbPassword string) {
	salt = EncodeMD5(stringHelper.RandString(10))
	dbPassword = EncodeMD5(password + salt)
	return
}

// HashSaltPassword 密码加密
func HashSaltPassword(salt string, password string) (dbPassword string) {
	dbPassword = EncodeMD5(password + salt)
	return
}

// CheckPasswordHash 验证密码
func CheckPasswordHash(password, salt, originPassword string) bool {
	return originPassword == EncodeMD5(password+salt)
}
