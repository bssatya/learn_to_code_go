package main

import (
	"fmt"
)

func main () {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)
	xi = append(xi[:3], xi[6:]...)
	fmt.Println(xi)	
}