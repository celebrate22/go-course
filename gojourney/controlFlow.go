package main

import (
	"errors"
	"fmt"
)

// FizzBuzz prints numbers 1..n, replacing multiples of 3 with "Fizz",
// multiples of 5 with "Buzz", and multiples of both with "FizzBuzz".
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// IsPrime reports whether n is a prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Divide returns a/b, or an error if b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("--- FizzBuzz(15) ---")
	FizzBuzz(15)

	fmt.Println("\n--- IsPrime ---")
	for _, n := range []int{1, 2, 17, 18, 29} {
		fmt.Printf("%d prime? %v\n", n, IsPrime(n))
	}

	fmt.Println("\n--- Divide ---")
	if result, err := Divide(10, 2); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}
	if _, err := Divide(5, 0); err != nil {
		fmt.Println("Error:", err)
	}
}
