package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)
	xi = append(xi, 52)
	fmt.Println(xi)
	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	xt := []int{56, 57, 58, 59, 60}
	xi = append(xi, xt...)
	fmt.Println(xi)

}
