package main

import (
	"fmt"
	"testing"
)

func Test_adder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("Expected '%v', but got '%v'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
