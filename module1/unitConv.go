package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// os.Args[0] is the program name. We need index 1 (value) and index 2 (unit).
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <value> <target_unit: C or F>")
		return
	}

	valStr := os.Args[1]
	// Normalizing the input unit to uppercase so 'c' and 'f' also work
	unit := strings.ToUpper(os.Args[2])

	// Converting the string input into a float64
	// The second argument (64) tells Go to parse it with float64 precision
	value, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		fmt.Printf("Error: '%s' is not a valid number.\n", valStr)
		return
	}

	// 2. Performing the conversion using a switch statement
	switch unit {
	case "C":
		// User provided Fahrenheit and wants to convert to Celsius
		celsius := (value - 32) * 5 / 9
		fmt.Printf("%.2f°F is equal to %.2f°C\n", value, celsius)

	case "F":
		// User provided Celsius and wants to convert to Fahrenheit
		fahrenheit := (value * 9 / 5) + 32
		fmt.Printf("%.2f°C is equal to %.2f°F\n", value, fahrenheit)

	default:
		fmt.Println("Error: Invalid target unit. Please use 'C' for Celsius or 'F' for Fahrenheit.")
	}
}
