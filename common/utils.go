package common

import (
	"reflect"
	"strings"
)

func Struct2Map(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := strings.ReplaceAll(v.Field(i).Tag.Get("json"), ",omitempty", "")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = Struct2Map(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
