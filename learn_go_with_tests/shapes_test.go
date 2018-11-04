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
			t.Errorf("Got (%f), Want (%f)", got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectange{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triange{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
