package env

import (
	"github.com/subosito/gotenv"
	"log"
	"os"
)

type Env struct {
	ApplicationToken string
	Environment      string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	DatabasePort     string
	BackendPort      string
}

func NewEnv() *Env {
	err := gotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	return &Env{
		ApplicationToken: os.Getenv("APPLICATION_TOKEN"),
		Environment:      os.Getenv("ENVIRONMENT"),
		DatabaseUser:     os.Getenv("POSTGRES_USER"),
		DatabasePassword: os.Getenv("POSTGRES_PASSWORD"),
		DatabaseName:     os.Getenv("POSTGRES_DB"),
		DatabaseHost:     os.Getenv("POSTGRES_HOST"),
		DatabasePort:     os.Getenv("POSTGRES_PORT"),
		BackendPort:      os.Getenv("BACKEND_PORT"),
	}
}
