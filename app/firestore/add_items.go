package firestore

import (
	"context"
	"log"
	"time"
)

func AddItems() {
	ctx := context.Background()

	client, err := GetClient()
	if err != nil {
		log.Fatalf("Failed to get Firestore client: %v", err)
	}
	defer client.Close()

	card := Cards{
		UserID:    "user123",
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Phone:     "123-456-7890",
		CardPath:  "/path/to/card",
		VoiceID:   []int64{1, 2, 3},
		CreatedAt: time.Now(),
	}

	_, _, err = client.Collection("cards").Add(ctx, card)
	if err != nil {
		log.Fatalf("Failed adding card: %v", err)
	} else {
		log.Println("Success adding card")
	}
}
