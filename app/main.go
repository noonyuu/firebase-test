package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()
	createSample(ctx, client)
}

func createClient(ctx context.Context) *firestore.Client {
	projectID := "test-project"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func createSample(ctx context.Context, client *firestore.Client) {
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}else{
		log.Println("Success adding alovelace")
	}
}
