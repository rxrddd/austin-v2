package arrayHelper

import "github.com/spf13/cast"

func ArrayStringIn(list []string, found string) bool {
	for _, s := range list {
		if found == s {
			return true
		}
	}
	return false
}
func ArrayInt64In(list []int64, found int64) bool {
	for _, s := range list {
		if found == s {
			return true
		}
	}
	return false
}

func ArrayAny2String(list []interface{}) []string {
	res := make([]string, 0)
	for _, s := range list {
		res = append(res, cast.ToString(s))
	}
	return res
}
