package main

import (
	"fmt"
	"math"
)

// Shape is an interface with an Area method
type Shape interface {
	Area() float64
}

// Rectangle holds width and height
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle holds radius
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	// Put both shapes in one slice of type Shape
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Circle{Radius: 3},
	}

	// Print every area from one loop
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
