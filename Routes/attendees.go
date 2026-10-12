package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "goapi.com/event/Models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	err = event.RegisterUserForEvent(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Attendee register successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.PostEvent
	event.ID = eventId

	err = event.CancelUserRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could cancel user registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Attendee registration Cancel!"})
}
