package main

import (
	"example.com/REST-API/DB"
	"example.com/REST-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
