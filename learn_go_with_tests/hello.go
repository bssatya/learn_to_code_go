package main

import "fmt"

const helloPrefixEng = "Hello, "
const helloPrevisSpanish = "Hola, "

func main() {
	fmt.Println(hello("Satya", "Spanish"))
}

func hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	var greetPrefix string
	if lang == "Spanish" {
		greetPrefix = helloPrevisSpanish
	} else {
		greetPrefix = helloPrefixEng
	}
	return greetPrefix + name
}
