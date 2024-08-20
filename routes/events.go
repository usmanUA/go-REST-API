package routes

import (
	"net/http"
	"strconv"

	"example.com/REST-API/models"
	"github.com/gin-gonic/gin"
)

func getEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event."})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, please try again later."})
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	userID := ctx.GetInt64("userID")
	event.UserID = userID
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events, please try again later."})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event."})
		return
	}
	userID := ctx.GetInt64("userID")
	if event.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update event."})
		return
	}
	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update events."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}
	userID := ctx.GetInt64("userID")
	event, err := models.GetEventByID(eventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event to delete."})
		return
	}
	if event.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete event."})
		return
	}
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update events."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
