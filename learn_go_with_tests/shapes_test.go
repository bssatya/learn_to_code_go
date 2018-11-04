package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got (%v), Want (%v)", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("Got (%v), Want (%v)", got, want)
	}
}
