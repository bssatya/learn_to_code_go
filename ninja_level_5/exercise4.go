package main

import (
	"fmt"
)

func main(){
	s := struct {
		first string
		friend map[string]int
		favDrinks []string
	} {
		first : "Raj",
		friend : map[string]int {
			"Vishnu" : 1223,
			"Ambi" : 12121,
		},
		favDrinks : []string{"filter kapi", "milk", "tender coconut"},
	}

	fmt.Println(s)
	fmt.Println(s.first)
	fmt.Println(s.friend)
	fmt.Println(s.favDrinks)

	for k, v := range s.friend {
		fmt.Println(k, v)
	}

	for i, v := range s.favDrinks {
		fmt.Println(i, v)
	}
}