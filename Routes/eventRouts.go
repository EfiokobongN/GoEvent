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
	authenticated.POST("/events/:id/attendees", registerForEvent)
	authenticated.DELETE("/events/:id/cancel", cancelRegistration)

	server.POST("/sign-up", signUp)
	server.POST("/login", login)
}
