package routes

import (
	_bookRepo "bookMg/business/book/repository/mysql"
	_bookUCase "bookMg/business/book/usecase"
	_bookHttp "bookMg/business/book/web"

	db "bookMg/db"

	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	api := router.Group("")
	bookRepo := _bookRepo.NewBookRepo(db.Client)
	bookUC := _bookUCase.NewBookUCase(bookRepo)
	_bookHttp.Init(api, bookUC)
}
