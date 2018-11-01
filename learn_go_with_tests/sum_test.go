package main

import "testing"

func Test_Sum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("Sum '%d', Expected '%d'", sum, expected)
	}
}
