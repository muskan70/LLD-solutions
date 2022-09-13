package main

import (
	_bookRepo "bookMg/business/book/repository/mysql"
	_bookUCase "bookMg/business/book/usecase"
	_bookHttp "bookMg/business/book/web"

	db "bookMg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("")
	bookRepo := _bookRepo.NewBookRepo(db.Client)
	bookUC := _bookUCase.NewBookUCase(bookRepo)
	_bookHttp.Init(api, bookUC)
	router.Run("localhost:8080")
}
