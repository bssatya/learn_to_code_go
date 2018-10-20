package main

import (
	"fmt"
)

func main() {
	f := getFunc()
	s := f()
	fmt.Println(s)
}

func getFunc() func() string {
	return func() string {
		return "Hello there!!"
	}
}