package main

import (
	"fmt"
)

func main () {
	x := "Ambi"

	if (x == "Raj") {
		fmt.Printf("Yes I am Raj")
	} else if (x == "Ambi") {
		fmt.Printf("Yes, nane Ambi")
	} else {
		fmt.Printf("Unknown actor")
	}
}