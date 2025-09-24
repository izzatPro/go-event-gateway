package main

import (
	"github.com/gin-gonic/gin"
	"github.com/izzatPro/go-event-gateway/db"
	"github.com/izzatPro/go-event-gateway/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")
}
