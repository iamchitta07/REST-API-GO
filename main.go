package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamchitta07/db"
	"github.com/iamchitta07/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
