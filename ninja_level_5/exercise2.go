package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	flavors []string
}


func main(){
	p1 := person {
		first : "Satya",
		last : "BS",
		flavors : []string{"tender coconut", "jackfruit"},
	}

	p2 := person {
		first : "Avinav",
		last : "Banuvalli",
		flavors : []string{"jackfruit", "tender cocounut"},
	}

	pm := map[string]person {
		p1.last : p1,
		p2.last : p2,
	}

	for k, v := range pm {
		fmt.Println("Key:", k)
		fmt.Println("Value:")
		fmt.Println("\t", v.first)
		fmt.Println("\t", v.last)
		for i, val := range v.flavors {
			fmt.Println ("\t\tAt Flavor Index", i, "Flavor", val)
		}
	}
}