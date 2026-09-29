package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "goapi.com/event/Models"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvent()
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

	eventSave.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created successfully"})
}
