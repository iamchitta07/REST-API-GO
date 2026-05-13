package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamchitta07/models"
	"github.com/iamchitta07/db"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", getHome)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data.", "error": err})
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
