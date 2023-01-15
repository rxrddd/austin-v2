package emptyHelper

import "github.com/spf13/cast"

func IsNotEmpty(val interface{}) bool {
	switch v := val.(type) {
	case string:
		return v != ""
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return cast.ToInt64(v) != 0
	}
	return false
}
