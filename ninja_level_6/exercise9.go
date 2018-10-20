package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	add(printSum, xi...)
}

func add(printFunc func(sum int), i ...int) {
	sum := 0
	for _, v := range i {
		sum += v
	}

	printFunc(sum)
}

func printSum(sum int) {
	fmt.Println("Sum is", sum)
}