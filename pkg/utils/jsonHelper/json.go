package jsonHelper

import "encoding/json"

func MustToString(value interface{}) string {
	marshal, _ := json.Marshal(value)
	return string(marshal)
}

func AnyToPtr(form interface{}, to interface{}) {
	_ = json.Unmarshal([]byte(MustToString(form)), &to)
}
