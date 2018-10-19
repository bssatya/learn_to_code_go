package main

import (
	"fmt"
)

func main () {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%v\t%#U\n", i, i)
		}
	}
}