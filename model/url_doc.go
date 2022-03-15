package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URLDoc struct {
	ID        primitive.ObjectID `bson:"id"`
	URLCode   string             `bson:"urlCode"`
	LongURL   string             `bson:"longURL"`
	ShortURL  string             `bson:"shortURL"`
	CreatedAt time.Time          `bson:"createdAt"`
	Expires   time.Time          `bson:"expires"`
}
