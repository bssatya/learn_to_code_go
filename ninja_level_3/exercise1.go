package main

import (
	"fmt"
)

func main () {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	
	i := 1
	for {
		if (i <= 10000) {
			break
		}
		fmt.Println(i)
	}

	for ;i <= 10000; {
		fmt.Println(i)
		i++
	}

}