package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area of shape is", s.area())
}
func main() {
	s := square {
		length : 4.0,
	}

	c := circle {
		radius : 3.0,
	}
	info(s)
	info(c)
}