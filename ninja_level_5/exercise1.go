package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	icFlavors []string
}

func main () {
	person1 := person {
		firstName : "Satya",
		lastName : "Banuvalli",
		icFlavors : []string{"tender coconut", "jackfruit", "mango"},
	}

	person2 := person {
		firstName : "Avinav",
		lastName : "Banuvalli",
		icFlavors : []string{"jackfruit", "tender coconut"},
	}

	fmt.Println(person1, "\t", person2)
	fmt.Printf("%T\n", person1)
}