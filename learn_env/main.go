package main

import (
	"log"
	"playground/learn_env/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	data := infrastructure.NewDriver()
	log.Println(data)
}
