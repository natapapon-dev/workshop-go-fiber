package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitENV() {
	log.Print("START LOAD ENV")
	err := godotenv.Load()
	log.Print("LOAD ENV COMPLETED")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
