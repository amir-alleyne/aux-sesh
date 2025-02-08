package main

import (
	"log"
	"os"

	"github.com/amir-alleyne/aux-sesh/backend/api/auth"
	"github.com/amir-alleyne/aux-sesh/backend/middleware"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"

	"github.com/clerk/clerk-sdk-go/v2"

	echo "github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	clerkKey := os.Getenv("CLERK_SECRET_KEY")

	clerk.SetKey(clerkKey)

	if _, err := auth.SetAuth(); err != nil {
		log.Fatal("Error setting auth:", err)
	}

	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	middleware.RegisterRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
