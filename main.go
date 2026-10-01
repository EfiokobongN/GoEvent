package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "goapi.com/event/DB"
	models "goapi.com/event/Models"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetche events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event fetched successfully", "events": events})
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
