package request

type GetBookById struct {
	Id int `form:"id" binding:"required"`
}
type DeleteBookById struct {
	Id int `form:"id" binding:"required"`
}
