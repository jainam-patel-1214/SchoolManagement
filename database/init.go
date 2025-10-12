package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitDb() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHostdsn := os.Getenv("DSN")

	fmt.Println("DB Host:", dbHostdsn)
	return dbHostdsn
}
