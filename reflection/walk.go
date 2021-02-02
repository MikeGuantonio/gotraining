package main

import "reflect"

func walk(x interface{}, fun func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		fun(val.Field(i).String())
	}
}