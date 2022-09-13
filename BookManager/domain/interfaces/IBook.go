package interfaces

import (
	"context"
	"dotpe/demo/domain/entity"
)

type IBookUcase interface {
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id int) (entity.Book, error)
	DeleteBookById(ctx context.Context, id int) error
	AddNewBook(ctx context.Context, req entity.Book) error
	UpdateDetails(ctx context.Context, req entity.Book) error
}
type IBookRepo interface {
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GetBookById(ctx context.Context, id int) (entity.Book, error)
	DeleteBookById(ctx context.Context, id int) error
	AddNewBook(ctx context.Context, req entity.Book) error
	UpdateDetails(ctx context.Context, req entity.Book) error
}
