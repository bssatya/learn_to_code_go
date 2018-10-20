package main

import (
	"fmt"
)

type person struct{
	first string
	last string
	age int
}
func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "and i am", p.age, "years old")
}

func main() {
	p1 := person {
		first : "Raj",
		last : "Kumar",
		age : 65,
	}
	p1.speak()
}