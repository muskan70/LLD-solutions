package main

import (
	"dotpe/demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(router)
	router.Run("localhost:8080")
}
