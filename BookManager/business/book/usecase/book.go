package usecase

import (
	"context"
	"dotpe/demo/domain/entity"
	"dotpe/demo/domain/interfaces"
)

type bookUCase struct {
	bookRepo interfaces.IBookRepo
}

func NewBookUCase(bookRepo interfaces.IBookRepo) interfaces.IBookUcase {
	return &bookUCase{
		bookRepo: bookRepo,
	}
}
func (buc *bookUCase) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	return buc.bookRepo.GetAllBooks(ctx)
}
func (buc *bookUCase) GetBookById(ctx context.Context, id int) (entity.Book, error) {
	return buc.bookRepo.GetBookById(ctx, id)
}
func (buc *bookUCase) DeleteBookById(ctx context.Context, id int) error {
	return buc.bookRepo.DeleteBookById(ctx, id)
}
func (buc *bookUCase) AddNewBook(ctx context.Context, req entity.Book) error {
	return buc.bookRepo.AddNewBook(ctx, req)
}
func (buc *bookUCase) UpdateDetails(ctx context.Context, req entity.Book) error {
	return buc.bookRepo.UpdateDetails(ctx, req)
}
