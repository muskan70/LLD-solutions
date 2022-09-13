package web

import (
	"bookMg/domain/entity"
	"bookMg/domain/interfaces"
	"bookMg/domain/request"
	"bookMg/domain/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bookUC interfaces.IBookUcase

func Init(api *gin.RouterGroup, buc interfaces.IBookUcase) {
	bookUC = buc
	api.GET("/api/allbooks", getallbooks)
	api.GET("/api/bookbyid", getbookbyid)
	api.DELETE("/api/deletebook", deletebookbyid)
	api.POST("/api/addbook", addnewbook)
	api.PUT("/api/update", updatedetails)
}
func getallbooks(c *gin.Context) {
	book2, err := bookUC.GetAllBooks(c)
	if err != nil {
		resp := response.Fail(false, "Bad Request", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, book2)

}
func getbookbyid(c *gin.Context) {
	var req request.GetBookById
	if err := c.BindQuery(&req); err != nil {
		resp := response.Fail(false, "Not valid input parameter", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	bb2, err := bookUC.GetBookById(c, req.Id)
	if err != nil {
		resp := response.Fail(true, "Data Missing", err)
		c.JSON(http.StatusOK, resp)
		return
	}
	if err == nil {
		c.JSON(http.StatusOK, bb2)
	}
}
func deletebookbyid(c *gin.Context) {
	var req request.DeleteBookById
	if err := c.BindQuery(&req); err != nil {
		resp := response.Fail(false, "Not valid input parameter", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	err2 := bookUC.DeleteBookById(c, req.Id)
	if err2 != nil {
		resp := response.Fail(true, "Book ID not found", err2)
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := response.Success("Successfully Deleted")
	c.JSON(http.StatusOK, resp)
}

func addnewbook(c *gin.Context) {
	var req entity.Book
	err := c.BindJSON(&req)
	if err != nil {
		resp := response.Fail(false, "Invalid Payload", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	err2 := bookUC.AddNewBook(c, req)
	if err2 != nil {
		resp := response.Fail(true, "Not able to insert data", err2)
		c.JSON(http.StatusOK, resp)
		return
	}
	if err2 == nil {
		resp := response.Success("Successfully inserted")
		c.JSON(http.StatusOK, resp)
	}

}
func updatedetails(c *gin.Context) {
	var req entity.Book
	err := c.BindJSON(&req)
	if err != nil {
		resp := response.Fail(false, "Bad Request", err)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if req.BookID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "False",
			"message": "Book ID is important for Update.",
		})
		return
	}
	err2 := bookUC.UpdateDetails(c, req)
	if err2 != nil {
		resp := response.Fail(true, "Not able to update", err)
		c.JSON(http.StatusOK, resp)
		return
	}
	resp := response.Success("Successfully Updated")
	c.JSON(http.StatusOK, resp)
}
