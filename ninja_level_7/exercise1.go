package main

import (
	"fmt"
)

func main() {
	v := 10
	fmt.Println("Address of value v =", v, "is", &v)
}