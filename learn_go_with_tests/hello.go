package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
	fmt.Println(hello("Satya"))
}

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}
