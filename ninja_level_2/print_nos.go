package main

import (
	"fmt"
)

func main () {
	
	for i := 33; i < 123; i++ {
		fmt.Printf("%c\t%d\n", i, i)
	}
}