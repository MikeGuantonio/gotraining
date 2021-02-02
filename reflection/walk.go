package main

import "reflect"

func walk(x interface{}, fun func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fun(field.String())
		}
	}
}