package firestore

import (
	"time"
)

type Cards struct {
	UserID    string    `firestore:"userId"`
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	Phone     string    `firestore:"phone"`
	CardPath  string    `firestore:"cardPath"`
	VoiceID   []int64   `firestore:"voiceId"`
	CreatedAt time.Time `firestore:"createdAt"`
}

type voices struct {
	userId    string    `firestore:"userID"`
	text      string    `firestore:"text"`
	cardId    []int64   `firestore:"cardId"`
	createdAt time.Time `firestore:"createdAt"`
}
