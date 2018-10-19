package main

import (
	"fmt"
)


func main () {
	ms := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being Evil", "Ice creams", "Sunsets"},
	}
	fmt.Println(ms)
	fmt.Printf("%T\n", ms)
	fmt.Println(len(ms))

	for k, kv := range ms {
		fmt.Println("Key is", k, "Value is", kv)
		for i, v := range kv {
			fmt.Println("\tValue at index", i, "is", v)
		}
	}
}