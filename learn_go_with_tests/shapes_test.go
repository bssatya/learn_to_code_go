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
			t.Errorf("%#v Got (%f), Want (%f)", shape, got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectange{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triange{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
