package app

import "go.mongodb.org/mongo-driver/mongo"
import "note-api/internal/config"

type App struct{
	DB *mongo.Client
	Config *config.Config
}