package controllers

import (
	"event-backend/config"
	"event-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&event)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"data":    event})
}

func GetEvents(c *gin.Context) {
	var events []models.Event

	config.DB.Find(&events)
	c.JSON(http.StatusOK, gin.H{
		"message": "Events retrieved successfully",
		"data":    events})
}

func GetEventByID(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	eventData := config.DB.First(&event, id).Error
	if eventData != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Detail Event retrieved successfully",
		"data":    event})
}

func UpdateEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	eventData := config.DB.First(&event, id).Error
	if eventData != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&event).Updates(event)
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"data":    event})
}

func DeleteEvent(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	eventData := config.DB.First(&event, id).Error
	if eventData != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	config.DB.Unscoped().Delete(&event)
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully"})
}
