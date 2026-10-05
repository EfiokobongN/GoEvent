package routes

import (
	"github.com/gin-gonic/gin"
	"goapi.com/event/middlewares"
)

func RegisterRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/delete/:id", deleteEvent)

	server.POST("/sign-up", signUp)
	server.POST("/login", login)
}
