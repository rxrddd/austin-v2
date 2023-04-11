package convertHelper

import (
	"strconv"
	"strings"
)

func Int64ArrayToString(arr []int64, delimiter string) string {
	if len(arr) == 0 {
		return ""
	}
	tmp := []string{}
	for _, v := range arr {
		s := strconv.FormatInt(v, 10)
		tmp = append(tmp, s)
	}
	return strings.Join(tmp, delimiter)
}
