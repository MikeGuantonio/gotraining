package main

import "reflect"

func walk(x interface{}, fun func(input string)) {
	val := reflect.ValueOf(x)

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
