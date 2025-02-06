package main

import (
	"log"

	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/amir-alleyne/aux-sesh/backend/api/handlers"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = auth.GetAdmin()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()

	//TODO : Add middleware to check if the user is authenticated

	// Register the routes
	handlers.RegisterRoutes(e)

	// Start the server on port 8080
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
