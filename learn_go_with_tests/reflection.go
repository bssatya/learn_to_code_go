package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("I still cant believe, souch korea beat Genrmeny 2-0 to push them to the bottom of their group")
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
