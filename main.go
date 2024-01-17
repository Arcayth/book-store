package main

import (
	"log"

	"github.com/haythamchanouni/book-store/database"
	"github.com/haythamchanouni/book-store/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := initApp(); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	router := routes.SetupRouter()
	router.Run(":8080")
}

func initApp() error {
	if err := loadEnv(); err != nil {
		return err
	}
	if err := database.StartDB(); err != nil {
		return err
	}
	return nil
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	return nil
}
