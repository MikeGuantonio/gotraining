package main

import "reflect"

func getValue(input interface{}) reflect.Value {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, fun func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
        walk(value.Interface(), fun)
    }

	switch val.Kind() {
	case reflect.String:
		fun(val.String())
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fun)
		}
	}
}
