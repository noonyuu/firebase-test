package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

var (
	firestoreClient *firestore.Client
	isInit          bool
)

func Init() error {

	var projectID = "test-project"
	var err error

	ctx := context.Background()
	firestoreClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}
	// 初期化済みにする
	isInit = true
	
	return nil
}

func GetClient() (*firestore.Client, error) {
	if !isInit {
		log.Fatalf("Firestore client is not initialized.")
	}
	return firestoreClient, nil
}
