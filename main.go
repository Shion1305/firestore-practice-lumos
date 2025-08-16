package main

import (
	"fmt"
	"time"
)

type (
	User struct {
		Name      string
		Age       int
		Email     string
		CreatedAt time.Time
	}
)

func main() {
	users := []User{
		{
			Name:      "John",
			Age:       20,
			Email:     "john@example.com",
			CreatedAt: time.Now(),
		},
		{
			Name:      "Sammy",
			Age:       21,
			Email:     "sammy@example.com",
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
	}

	fmt.Println(users)
}
