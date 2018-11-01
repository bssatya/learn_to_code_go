package main

import "testing"

func Test_Repeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got '%s', Want '%s'", repeated, expected)
	}
}