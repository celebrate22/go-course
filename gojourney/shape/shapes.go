package shapes

import (
	"fmt"
	"math"
)

// ErrInvalidDimension is a custom error for negative or zero values.
type ErrInvalidDimension struct {
	ShapeType string
	Message   string
}

func (e *ErrInvalidDimension) Error() string {
	return fmt.Sprintf("[%s error]: %s", e.ShapeType, e.Message)
}

// Shape is an interface with an Area method
type Shape interface {
	Area() float64
}

// Rectangle holds width and height
type Rectangle struct {
	Width, Height float64
}

// NewRectangle validates dimensions and returns a custom error if invalid
func NewRectangle(width, height float64) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, &ErrInvalidDimension{
			ShapeType: "Rectangle",
			Message:   "width and height must be greater than zero",
		}
	}
	return Rectangle{Width: width, Height: height}, nil
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle holds radius
type Circle struct {
	Radius float64
}

// NewCircle validates radius and returns a custom error if invalid
func NewCircle(radius float64) (Circle, error) {
	if radius <= 0 {
		return Circle{}, &ErrInvalidDimension{
			ShapeType: "Circle",
			Message:   "radius must be greater than zero",
		}
	}
	return Circle{Radius: radius}, nil
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
