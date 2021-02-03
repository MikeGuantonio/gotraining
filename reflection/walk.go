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

	switch val.Kind() {
	case reflect.String:
		fun(val.String())
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fun)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), fun)
		}
	}
	
}
