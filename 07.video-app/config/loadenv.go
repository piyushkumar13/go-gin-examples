package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvs() {

	err := godotenv.Load("local.env")

	if err != nil {
		log.Fatal("Error loading local.env file", err)
	}
}
