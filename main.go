package main

import (
	"errors"
	"fmt"
	"gojourney/shape" // Replace with your actual module path
)

func main() {
	var shapeList []shapes.Shape

	// 1. Create a valid Rectangle
	rect, err := shapes.NewRectangle(4, 5)
	if err != nil {
		handleError(err)
	} else {
		shapeList = append(shapeList, rect)
	}

	// 2. Create a valid Circle
	circle, err := shapes.NewCircle(3)
	if err != nil {
		handleError(err)
	} else {
		shapeList = append(shapeList, circle)
	}

	// 3. Example: Attempting to create an invalid shape to trigger the custom error
	invalidRect, err := shapes.NewRectangle(-2, 5)
	if err != nil {
		handleError(err) // This will catch and print the custom error
	} else {
		shapeList = append(shapeList, invalidRect)
	}

	fmt.Println("\n--- Calculating Areas ---")
	for _, shape := range shapeList {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

// handleError intercepts custom shapes errors specifically
func handleError(err error) {
	var shapeErr *shapes.ErrInvalidDimension
	if errors.As(err, &shapeErr) {
		fmt.Printf("Caught expected custom validation error -> Type: %s | Message: %s\n",
			shapeErr.ShapeType, shapeErr.Message)
	} else {
		fmt.Printf("Standard Error: %v\n", err)
	}
}
