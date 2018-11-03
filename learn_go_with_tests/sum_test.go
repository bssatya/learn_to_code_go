package main

import (
	"reflect"
	"testing"
)

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

func Test_SumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	expected := []int{6, 15, 24}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got (%v), Expected (%v)", got, expected)
	}
}
