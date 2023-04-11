package model

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type JsonSlice []string //json数组数据

type SplitSlice []string //逗号隔开的数据

func (s *JsonSlice) Scan(v interface{}) error {
	if vt, ok := v.([]byte); ok {
		return json.Unmarshal(vt, &s)
	}
	return nil
}

func (s JsonSlice) Value() (driver.Value, error) {
	m, _ := json.Marshal(s)
	return string(m), nil
}

func (s *SplitSlice) Scan(v interface{}) error {
	if vt, ok := v.([]byte); ok {
		*s = strings.Split(string(vt), ",")
	}
	return nil
}

func (s SplitSlice) Value() (driver.Value, error) {
	m := strings.Join(s, ",")
	return m, nil
}
