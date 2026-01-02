package config

import(
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	MongodRI string
}

func LoadConfig() *Config{
	 
    err:= godotenv.Load()
    
	if err!=nil{
       log.Println("No .env found")
	}

	port:= os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
    
	mongoUri := os.Getenv("MONGO_URI")

	if mongoUri == "" {
		log.Fatal("Mongo URI not found")
	}

	return &Config{
		Port: port,
		MongodRI: mongoUri,
	}
}