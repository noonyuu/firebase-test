package main

import (
	// "context"
	// "flag"
	// "fmt"
	// "log"

	// "google.golang.org/api/iterator"

	// "cloud.google.com/go/firestore"

	"log"

	"github.com/noonyuu/firebase-test/app/firestore"
)

func main() {
	err := firestore.Init()
	if err != nil {
		log.Fatalln(err)
	}

	firestore.AddItems()
}
