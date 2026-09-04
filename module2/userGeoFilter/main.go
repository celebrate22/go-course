package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
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

	// 2. Filter: non-empty email AND Address.Country == "Nigeria"
	filtered := []User{}
	for _, u := range users {
		if u.Email != "" && u.Address.Country == "Nigeria" {
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