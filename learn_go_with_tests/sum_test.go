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

func Test_SumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got (%v), Expected (%v)", got, expected)
		}
	}

	t.Run("make sum of tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
		expected := []int{5, 11, 17}

		checkSums(t, got, expected)
	})

	t.Run("check safely sum up empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3}, []int{4, 5, 6})
		expected := []int{0, 5, 11}

		checkSums(t, got, expected)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	}
}
