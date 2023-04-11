package typeConvert

func ClearMapZeroValue(params map[string]interface{}) map[string]interface{} {
	for k, v := range params {
		needDelete := false
		// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, byte
		switch v.(type) {
		case int:
			if v.(int) == 0 {
				needDelete = true
			}
		case int8:
			if v.(int8) == 0 {
				needDelete = true
			}
		case int16:
			if v.(int16) == 0 {
				needDelete = true
			}
		case int32:
			if v.(int32) == 0 {
				needDelete = true
			}
		case int64:
			if v.(int64) == 0 {
				needDelete = true
			}
		case uint:
			if v.(uint) == 0 {
				needDelete = true
			}
		case uint8:
			if v.(uint8) == 0 {
				needDelete = true
			}
		case uint16:
			if v.(uint16) == 0 {
				needDelete = true
			}
		case uint32:
			if v.(uint32) == 0 {
				needDelete = true
			}
		case uint64:
			if v.(uint64) == 0 {
				needDelete = true
			}
		case float32:
			if v.(float32) == 0 {
				needDelete = true
			}
		case float64:
			if v.(float64) == 0 {
				needDelete = true
			}
		case string:
			if v.(string) == "" {
				needDelete = true
			}
		case bool:
			if v.(bool) == false {
				needDelete = true
			}
		case interface{}:
			if v.(interface{}) == nil {
				needDelete = true
			}
		}
		if needDelete {
			delete(params, k)
		}
	}
	return params
}
