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

	numberOfValues := 0
	var getField func (int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fun(val.String())
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fun)
	}
	
}
