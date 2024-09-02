package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id"})
		return
	}

	event, err := models.QueryEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not query event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not register for event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered for event"})
}

func unregisterForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse event id"})
		return
	}

	event, err := models.QueryEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not query event"})
		return
	}

	err = event.Unregister(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not unregister for event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "unregistered for event"})
}
