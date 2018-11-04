package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectange := Rectange{10.0, 10.0}
	got := Perimeter(rectange)
	want := 40.0

	if got != want {
		t.Errorf("Got (%0.2f), Want (%0.2f)", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("Got (%0.2f), Want (%0.2f)", got, want)
		}
	}

	t.Run("Area of Rectange", func(t *testing.T) {
		rectangle := Rectange{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("Area of Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.16)
	})
}
