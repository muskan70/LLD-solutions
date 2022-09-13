package routes

import (
	"dotpe/demo/bootstrap"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("")

	bootstrap.Init(api)
}
