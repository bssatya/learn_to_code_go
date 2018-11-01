package main

import (
	"fmt"
)

// func Sum takes five integers and returns the sume of all the 5 integers
func Sum(numbers [5]int) (sum int) {
	for _, n := range numbers {
		fmt.Println(sum)
		sum += n
	}
	return
}
