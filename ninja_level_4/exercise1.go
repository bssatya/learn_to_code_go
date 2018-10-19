package main

import (
	"fmt"
)

func main () {

	ai := [5]int{1, 2, 4, 5, 6}

	for i := 0; i < len(ai); i++ {
		ai[i] = 50 + i
	}

	fmt.Printf("%T\n", ai)
	for i, v := range ai {
		fmt.Println(i, v)
	}
}
