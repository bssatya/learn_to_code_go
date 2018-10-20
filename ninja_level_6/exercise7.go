package main

import (
	"fmt"
)

var x int
var y func()

func main() {
	f := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	f()
	y = f
	y()
}