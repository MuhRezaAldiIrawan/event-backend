package main

import (
	"event-backend/config"
	"event-backend/models"
	"log"
	"net/http"

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
		api.GET("/events", getEvents)
		api.POST("/events", createEvents)
	}

	// Start the server on port 8080
	server.Run(":8080")
}

func mainHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is running!",
	})
}

func getEvents(c *gin.Context) {
	events := models.ShowEvents()

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvents(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request body",
			"error":   err.Error(),
		})
		return
	}

	// Save the event to the in-memory slice
	event.AddEvent()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})

}
