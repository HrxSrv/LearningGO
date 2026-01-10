package domain

import "errors"

var (
	MongoDocNotFound = errors.New("Doc not found in mongodb")
	ErrNotFound = errors.New("Doc Not found")
)