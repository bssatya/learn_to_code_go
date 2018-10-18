package main

import (
	"fmt"
)

func main () {
	a := 42 == 42
	b := 42 <= 42
	c := 42 >= 42
	d := 42 != 42
	e := 42 < 42
	f := 42 > 42

	fmt.Printf("42 == 42", a)
	fmt.Printf("\n42 <= 42", b)
	fmt.Printf("\n42 >= 42", c)
	fmt.Printf("\n42 != 42", d)
	fmt.Printf("\n42 < 42", e)
	fmt.Printf("\n42 > 42", f)
}