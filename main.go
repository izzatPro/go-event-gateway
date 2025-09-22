package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izzatPro/go-event-gateway/db"
	"github.com/izzatPro/go-event-gateway/models"
)

func main() {
	db.InitDB()

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// сервер проставляет служебные поля
	event.UserID = 1

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
