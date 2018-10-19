package main

import (
	"fmt"
)

func main () {
	by := 1980

	for {
		if by > 2018 {
			break
		}
		fmt.Println(by)
		by++
	}
}