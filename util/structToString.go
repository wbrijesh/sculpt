package util

import "reflect"

func StructToString(object interface{}) string {
	return reflect.TypeOf(object).Name()
}
