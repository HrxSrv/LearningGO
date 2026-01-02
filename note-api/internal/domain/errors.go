package domain

import "errors"

var (
	MongoDocNotFound = errors.New("Doc not found in mongodb")
)