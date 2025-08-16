package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

type (
	User struct {
		Name      string
		Age       int
		Email     string
		CreatedAt time.Time
	}
)

const (
	FirestoreProjectID  = "your-project-id"
	CollectionNameUsers = "users"
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

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, FirestoreProjectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	// Push users to Firestore "users" collection
	for i, user := range users {
		docRef := client.Collection(CollectionNameUsers).NewDoc()
		_, err := docRef.Set(ctx, user)
		if err != nil {
			log.Printf("Failed to add user %d to Firestore: %v", i, err)
			continue
		}
		fmt.Printf("User %s added to Firestore with ID: %s\n", user.Name, docRef.ID)
	}

	fmt.Println("All users processed")
}
