package main

import (
	"event-backend/config"
	"event-backend/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	server := gin.Default()

	// Define routes
	api := server.Group("/api/v1")
	{

		// event routes
		api.GET("/events", controllers.GetEvents)
		api.POST("/events", controllers.CreateEvent)
		api.GET("/events/:id", controllers.GetEventByID)
		api.PUT("/events/:id", controllers.UpdateEvent)
		api.DELETE("/events/:id", controllers.DeleteEvent)

		// user routes
		api.POST("/register", controllers.RegisterUser)
	}

	// Start the server on port 8080
	server.Run(":8080")
}
