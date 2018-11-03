package main

import "testing"

func Test_Sum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("Sum '%d', Expected '%d'", sum, expected)
		}
	})

}
