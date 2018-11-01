package main

import (
	"fmt"
	"testing"
)

func Test_Repeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got '%s', Want '%s'", repeated, expected)
	}
}
func ExampleRepeat() {
	sum := Repeat("s", 5)
	fmt.Println(sum)
	// Output: sssss
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
