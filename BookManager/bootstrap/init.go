package bootstrap

import (
	_bookRepo "dotpe/demo/business/book/repository/mysql"
	_bookUCase "dotpe/demo/business/book/usecase"
	_bookHttp "dotpe/demo/business/book/web"

	db "dotpe/demo/db"

	"github.com/gin-gonic/gin"
)

func Init(api *gin.RouterGroup) {
	bookRepo := _bookRepo.NewBookRepo(db.Client)
	bookUC := _bookUCase.NewBookUCase(bookRepo)
	_bookHttp.Init(api, bookUC)
}
