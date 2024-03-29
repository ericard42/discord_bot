package internal

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
