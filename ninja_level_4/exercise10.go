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

	ms["fleming_ian"] = []string{"Steaks", "Cigars", "Espionage"}

	printMap(ms)
	delete(ms, "no_dr")
	printMap(ms)
}

func printMap(ms map[string][]string) {
	for k, kv := range ms {
		fmt.Println("Key is", k, "Value is", kv)
		for i, v := range kv {
			fmt.Println("\tValue at index", i, "is", v)
		}
	}
}