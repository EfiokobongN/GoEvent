package main

import (
	"github.com/gin-gonic/gin"
	db "goapi.com/event/DB"
	routes "goapi.com/event/Routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoute(server)
	server.Run(":8080")
}
