package main

func walk(x interface{}, fun func(input string)) {
	fun("Nick")
}