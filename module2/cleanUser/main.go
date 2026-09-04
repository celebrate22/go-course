// Write a program that:

// 1. Reads a file `users.json` containing a JSON array of your `User`-like structs (reuse or adapt the struct from 2.1)
// 2. Filters to only users whose `Email` field is non-empty
// 3. Writes the filtered list to `filtered_users.json`, pretty-printed (hint: look at `json.MarshalIndent`)
// 4. If `users.json` doesn't exist, print a clear error instead of crashing
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// User mirrors the struct from 2.1 — adjust fields/tags as needed
// to match whatever you used there.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 1. Read users.json
	data, err := os.ReadFile("users.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Error: users.json not found in the current directory.")
			return
		}
		fmt.Println("Error reading users.json:", err)
		return
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Error parsing users.json:", err)
		return
	}

	// 2. Filter to users with non-empty Email
	// Initialized as []User{}, not var filtered []User — so if zero
	// users match, this still marshals to `[]`, not `null`.
	filtered := []User{}
	for _, u := range users {
		if u.Email != "" {
			filtered = append(filtered, u)
		}
	}

	// 3. Write pretty-printed JSON to filtered_users.json
	out, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		fmt.Println("Error encoding filtered users:", err)
		return
	}

	if err := os.WriteFile("filtered_users.json", out, 0644); err != nil {
		fmt.Println("Error writing filtered_users.json:", err)
		return
	}

	fmt.Printf("Wrote %d of %d users to filtered_users.json\n", len(filtered), len(users))
}