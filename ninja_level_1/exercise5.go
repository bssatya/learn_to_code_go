package main

import (
	"fmt"
)

type regNo int

var x regNo
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}
