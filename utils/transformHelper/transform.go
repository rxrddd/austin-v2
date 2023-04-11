package transformHelper

import "github.com/spf13/cast"

func ArrayStringToInt64(ids []string) []int64 {
	var newIds []int64
	for _, id := range ids {
		newIds = append(newIds, cast.ToInt64(id))
	}
	return newIds
}

func ArrayInt64ToString(ids []int64) []string {
	var newIds []string
	for _, id := range ids {
		newIds = append(newIds, cast.ToString(id))
	}
	return newIds
}
