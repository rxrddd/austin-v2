package makeHelper

import (
	"reflect"
	"strings"
)

type Rule func(string) string

var maskRule map[string]Rule

const (
	password = "password"
	mobile   = "mobile"
	email    = "email"
)

func init() {
	maskRule = make(map[string]Rule)
	RegisterMaskRule(password, func(s string) string {
		return "****"
	})
	RegisterMaskRule(mobile, func(s string) string {
		return s[:3] + "****" + s[7:]
	})
	RegisterMaskRule(email, func(s string) string {
		email := strings.Split(s, "@")
		emailPrefix := email[0]
		if len(email) == 1 {
			s = emailPrefix[:1] + strings.Repeat("*", len(emailPrefix)-1) + emailPrefix[len(emailPrefix)-1:]
		} else {
			s = emailPrefix[:1] + strings.Repeat("*", len(emailPrefix)-1) + emailPrefix[len(emailPrefix)-1:] + "@" + email[1]
		}
		return s
	})
}
func RegisterMaskRule(s string, rule Rule) {
	maskRule[s] = rule
}

func Mask(req interface{}) {
	elem := reflect.Indirect(reflect.ValueOf(req))
	for i := 0; i < elem.NumField(); i++ {
		valueField := elem.Field(i)
		typeField := elem.Type().Field(i)
		//如果字段不可写直接跳过 比如pb文件的 state,sizeCache,unknownFields等字段
		if !valueField.CanSet() {
			continue
		}
		switch valueField.Kind() {
		case reflect.String:
			//判断如果有mask tag标签才进行脱敏
			maskString(valueField, typeField)
		case reflect.Ptr:
			//如果是指针，判断不是nil 并且原值是struct 执行
			if !valueField.IsNil() && reflect.TypeOf(valueField.Elem().Interface()).Kind() == reflect.Struct {
				Mask(valueField.Interface())
			}
		case reflect.Struct:
			//如果是struct获取struct的指针
			Mask(valueField.Addr().Interface())
		case reflect.Slice, reflect.Array:
			reflectLen := valueField.Len()
			for rf := 0; rf < reflectLen; rf++ {
				//把切片每个数据的下标数据的指针传递进去
				rfKind := valueField.Index(rf).Kind()
				switch rfKind {
				case reflect.Ptr:
					Mask(valueField.Index(rf).Interface())
				case reflect.Struct:
					Mask(valueField.Index(rf).Addr().Interface())
				}
			}
		}
	}
}
func maskString(valueField reflect.Value, typeField reflect.StructField) {
	tag := typeField.Tag
	if value, ok := tag.Lookup("mask"); ok {
		if valueField.IsValid() && valueField.CanSet() {
			if fc, b := maskRule[value]; b {
				valueField.SetString(fc(valueField.String()))
			}
		}
	}
}
func maskEmail(s string) string {
	email := strings.Split(s, "@")
	emailPrefix := email[0]
	if len(email) == 1 {
		s = emailPrefix[:1] + strings.Repeat("*", len(emailPrefix)-1) + emailPrefix[len(emailPrefix)-1:]
	} else {
		s = emailPrefix[:1] + strings.Repeat("*", len(emailPrefix)-1) + emailPrefix[len(emailPrefix)-1:] + "@" + email[1]
	}
	return s
}
