package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamchitta07/models"
)

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event Id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later."})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later."})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getHome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to our Event Portal"})
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindBodyWithJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event Id"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later."})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data."})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event Updated successfully."})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Event Id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later."})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
