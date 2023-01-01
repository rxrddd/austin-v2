package encrypt

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
