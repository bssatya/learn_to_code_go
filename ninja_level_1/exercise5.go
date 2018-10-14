package main

import (
	"fmt"
)

type regNo int

var x regNo
var y regNo

func main () {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
}