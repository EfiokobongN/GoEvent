package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/event", getEvent)
	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Event fetched successfully"})
}
