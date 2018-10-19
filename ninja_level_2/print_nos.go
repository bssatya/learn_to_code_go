package main

import (
	"fmt"
)

func main () {
	
	for i := 33; i < 123; i++ {
		fmt.Printf("%#U\t%#x\t%v\n", i, i, i)
	}
}