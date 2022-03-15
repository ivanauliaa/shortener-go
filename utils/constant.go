package utils

import "context"

const (
	BASE_URL      = "http://localhost:5100/"
	PORT          = ":5100"
	DB_URI        = "mongodb://root:example@mongo:27017"
	DB_NAME       = "shortener_go"
	DB_COLLECTION = "urls"
)

var GLOBAL_CONTEXT = context.TODO()
