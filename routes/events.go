package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(c *gin.Context) {
	events, err := models.QueryEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id"})
		return
	}

	event, err := models.QueryEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse data"})
		return
	}

	userId := c.GetInt64("userId")
	event.UserID = userId
	fmt.Println(userId)

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.QueryEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you do not have permission to update this event"})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.QueryEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch event"})
		return
	}

	if userId != event.UserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you do not have permission to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted", "event": event})
}
