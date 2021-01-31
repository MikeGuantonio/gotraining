package main

import "reflect"

func walk(x interface{}, fun func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)

	fun(field.String())
}