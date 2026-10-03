package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	models "goapi.com/event/Models"
)

func signUp(context *gin.Context) {
	var userReg models.User
	err := context.ShouldBindJSON(&userReg)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = userReg.Save()
	if err != nil {
		log.Println("signup save error:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register account. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Account register successfully", "event": userReg})
}
