package main

import (
	"fmt"
)

type regNo int

var x regNo

func main () {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}