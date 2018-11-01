package main

import "strings"

// func Repeat takes a string and returns a string which repeated form of the input string
func Repeat(s string, times int) (repeat string) {
	repeat = strings.Repeat(s, times)
	return
}
