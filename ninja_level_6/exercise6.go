package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Inside ananymouse func")
	}()
}