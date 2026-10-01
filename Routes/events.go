package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "goapi.com/event/Models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event fetched successfully", "events": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var eventSave models.PostEvent
	err := context.ShouldBindJSON(&eventSave)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	eventSave.ID = 1
	eventSave.UserID = 6

	err = eventSave.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created successfully", "event": eventSave})
}
