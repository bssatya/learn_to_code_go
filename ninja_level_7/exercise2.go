package main

import (
	"fmt"
)

type person struct{
	first string
	last string
	age int
}
func main() {
	person := person{"Raj", "Kumar", 75}
	fmt.Println("Person Value in Main func before calling ChangeMe", person)
	changeMe(&person)
	fmt.Println("Person Value in Main func after calling ChangeMe", person)
}

func changeMe(person *person) {
	fmt.Println("Person Value in ChangeMe func before changing", *person)
	person.age = 65
	fmt.Println("Person Value in ChangeMe func after changing", *person)
}