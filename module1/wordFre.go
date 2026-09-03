package main

import (
	"fmt"
	"strings"
)

func main() {
	// Sample text containing repeating words and varied casing
	text := "Go is an open source programming language that makes it easy to build simple reliable and efficient software Go is great"

	// 1. Split the string into a slice of individual words
	// strings.Fields handles multiple spaces or tabs automatically
	words := strings.Fields(text)

	// 2. Initialize the map using the built-in make function
	// A map map[string]int uses strings as keys and integers as values
	wordCounts := make(map[string]int)

	// 3. Loop through the slice and increment counters
	for _, word := range words {
		// Normalize to lowercase so "Go" and "go" are counted together
		lowerWord := strings.ToLower(word)
		
		// If the key doesn't exist yet, Go defaults its value to 0, 
		// so we can safely increment it directly without a checking block.
		wordCounts[lowerWord]++
	}

	// 4. Print out the final word dashboard
	fmt.Println("--- Word Frequency Dashboard ---")
	for word, count := range wordCounts {
		fmt.Printf("%-12s : %d\n", word, count)
	}
}
