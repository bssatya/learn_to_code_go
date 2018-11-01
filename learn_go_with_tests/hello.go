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

func Hello(name string, lang string) (prefix string) {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = helloPrefixSpanish
	case french:
		prefix = helloPrefixFrench
	default:
		prefix = helloPrefixEng
	}

	return
}
