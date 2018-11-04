package main

import "math"

type Shape interface {
	Area() float64
}
type Rectange struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectange) Area() (area float64) {
	area = r.Height * r.Width
	return
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

func Perimeter(rectangle Rectange) (perimeter float64) {
	perimeter = 2*rectangle.Width + 2*rectangle.Height

	return
}

func Area(rectangle Rectange) (area float64) {
	area = rectangle.Height * rectangle.Width
	return
}
