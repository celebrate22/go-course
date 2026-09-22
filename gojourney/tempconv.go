package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toFahrenheit(c float64) float64 { return c*9/5 + 32 }
func toCelsius(f float64) float64    { return (f - 32) * 5 / 9 }
func toKelvin(c float64) float64     { return c + 273.15 }
func kelvinToCelsius(k float64) float64 { return k - 273.15 }

func celsiusToAll(c float64) (float64, float64, float64) {
	return c, toFahrenheit(c), toKelvin(c)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Temperature Converter")
	fmt.Println("Enter a value with its unit:)
	fmt.Print("> ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToUpper(input))

	if len(input) < 2 {
		fmt.Println("Invalid input.")
		return
	}

	unit := input[len(input)-1]
	valueStr := input[:len(input)-1]

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Could not parse number:", valueStr)
		return
	}

	var c, f, k float64

	switch unit {
	case 'C':
		c, f, k = celsiusToAll(value)
	case 'F':
		c = toCelsius(value)
		_, f, k = celsiusToAll(c)
	case 'K':
		c = kelvinToCelsius(value)
		_, f, _ = celsiusToAll(c)
		k = value
	default:
		fmt.Println("Unknown unit. Use C, F, or K.")
		return
	}

	fmt.Printf("\n%.2f°C = %.2f°F = %.2fK\n", c, f, k)
}
