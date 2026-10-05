package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "goapi.com/event/Util"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorize"})
		return
	}

	userId, err := util.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorize"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
