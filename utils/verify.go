package utils

import "reflect"

// IsAnyBlank 判断传入的值是否有空值，如果有为空，就返回 true
func IsAnyBlank(values ...interface{}) bool {
	for _, v := range values {
		if isBlank(v) {
			return true
		}
	}
	return false
}

// isBlank 判断单个值是否为空
func isBlank(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}
