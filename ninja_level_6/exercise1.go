package main

import (
	"fmt"
)


func main () {
	fmt.Println(foo())
	i, s := bar()
	fmt.Println(i, s)
}

func foo() int {
	return 2018
}

func bar() (int, string) {
	return 1980, "Birth Year"
}