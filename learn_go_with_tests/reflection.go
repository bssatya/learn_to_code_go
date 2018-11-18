package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("I still cant believe, souch korea beat Genrmeny 2-0 to push them to the bottom of their group")
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}
