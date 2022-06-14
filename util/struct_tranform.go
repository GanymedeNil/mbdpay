package util

import (
	"fmt"
	"reflect"
	"strings"
)

// ToMap 结构体转为Map[string]interface{}
func ToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			val := strings.Split(tagValue, ",")
			if v.Field(i).Kind() == reflect.Pointer {
				if !v.Field(i).IsNil() {
					out[val[0]] = v.Field(i).Elem().Interface()
				}
			} else {
				out[val[0]] = v.Field(i).Interface()
			}
		}
	}
	return out, nil
}
