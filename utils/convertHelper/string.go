package convertHelper

import (
	"strconv"
	"strings"
)

func StringToInt64Array(str string, delimiter string) (res []int64, err error) {
	if str == "" {
		return
	}
	tmp := strings.Split(str, delimiter)
	for _, v := range tmp {
		if value, err := strconv.ParseInt(v, 10, 64); err != nil {
			return []int64{}, err
		} else {
			res = append(res, value)
		}

	}
	return
}

func StringToInt64ArrayNoErr(str string, delimiter string) (res []int64) {
	if str == "" {
		return
	}
	tmp := strings.Split(str, delimiter)
	for _, v := range tmp {
		if value, err := strconv.ParseInt(v, 10, 64); err != nil {
			return []int64{}
		} else {
			res = append(res, value)
		}

	}
	return
}
