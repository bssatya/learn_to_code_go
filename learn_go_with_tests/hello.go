package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefixEng = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func main() {
	fmt.Println(hello("Satya", "Spanish"))
}

func hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	var greetPrefix string
	if lang == spanish {
		greetPrefix = helloPrefixSpanish
	} else if lang == french {
		greetPrefix = helloPrefixFrench
	} else {
		greetPrefix = helloPrefixEng
	}
	return greetPrefix + name
}
