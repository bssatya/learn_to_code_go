package main

import (
	"fmt"
)

func main () {
	by := 1980
	cy := 2018
	for cy >= by {
		fmt.Println(by)
		by++		
	}
}