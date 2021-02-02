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

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fun(field.String())
		case reflect.Struct:
			walk(field.Interface(), fun)
		}
	}
}
