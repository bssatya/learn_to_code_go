package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("I still cant believe, souch korea beat Genrmeny 2-0 to push them to the bottom of their group")
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
