package main

import (
	"fmt"
)

func main(){
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Foo Sum =", foo(xi...))
	fmt.Println("Bar Sum =", bar(xi))
}

func foo(i ...int) int {
	sum := 0;
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar (xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
