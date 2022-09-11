package main

import (
	"log"

	"github.com/YogaPratama02/go-crud-mongo/routes"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}
	routes.Init()
}
