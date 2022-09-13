package mysql

import (
	"context"
	"database/sql"
	"dotpe/demo/db"
	"dotpe/demo/domain/entity"
	"dotpe/demo/domain/interfaces"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type bookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) interfaces.IBookRepo {
	return &bookRepo{db: db}
}

func (brepo *bookRepo) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	var book2 []entity.Book
	db1, err := db.Init()
	if err != nil {

		return nil, err
	}
	defer db1.Close()
	queryBuilder := sq.Select("ID", "book_name", "author", "price", "quantity", "sold", "in_stock", "status")
	queryBuilder = queryBuilder.From("book")
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := db1.QueryContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var bb entity.Book
		if err := rows.Scan(&bb.BookID, &bb.Name, &bb.Writer, &bb.Price, &bb.Quantity, &bb.Sold, &bb.InStock, &bb.Status); err != nil {
			log.Println(err)
			return nil, err
		}
		book2 = append(book2, bb)
	}

	return book2, nil

}
func (brepo *bookRepo) GetBookById(ctx context.Context, id int) (entity.Book, error) {
	fmt.Println(id)
	var bb entity.Book
	db1, err := db.Init()
	if err != nil {
		log.Println(err)
		return bb, err
	}
	defer db1.Close()
	queryBuilder := sq.Select("ID", "book_name", "author", "price", "quantity", "sold", "in_stock", "status")
	queryBuilder = queryBuilder.From("book").Where(sq.Eq{"ID": id})
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return bb, err
	}
	rows := db1.QueryRow(query, qargs...)
	if err = rows.Scan(&bb.BookID, &bb.Name, &bb.Writer, &bb.Price, &bb.Quantity, &bb.Sold, &bb.InStock, &bb.Status); err != nil {
		log.Println(err)
		return bb, err
	}
	return bb, nil

}
func (brepo *bookRepo) DeleteBookById(ctx context.Context, id int) error {
	db1, err := db.Init()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db1.Close()
	queryBuilder := sq.Delete("book").Where(sq.Eq{"ID": id})
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	row, err := db1.QueryContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err)
		return err
	}

	defer row.Close()
	_, err = db1.ExecContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
func (brepo *bookRepo) AddNewBook(ctx context.Context, req entity.Book) error {
	db1, err := db.Init()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db1.Close()
	queryBuilder := sq.Insert("book").Columns("ID", "book_name", "author", "price", "quantity", "sold", "in_stock", "status").Values(req.BookID, req.Name, req.Writer, req.Price, req.Quantity, 0, req.Quantity, "A")
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err2 := db1.QueryContext(ctx, query, qargs...)
	if err2 != nil {
		log.Println(err)
		return err2
	}

	_, err3 := db1.ExecContext(ctx, query, qargs...)
	if err2 != nil {
		fmt.Println(err3)
		return err3
	}
	return nil

}
func (brepo *bookRepo) UpdateDetails(ctx context.Context, req entity.Book) error {
	db1, err := db.Init()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db1.Close()
	queryBuilder := sq.Update("book")
	if req.Name != "" {
		queryBuilder = queryBuilder.Set("book_name", req.Name).Where(sq.Eq{"ID": req.BookID})
	}
	if req.Writer != "" {
		queryBuilder = queryBuilder.Set("author", req.Writer).Where(sq.Eq{"ID": req.BookID})
	}
	if req.Price != 0 {
		queryBuilder = queryBuilder.Set("price", req.Price).Where(sq.Eq{"ID": req.BookID})
	}
	if req.Quantity != 0 {
		queryBuilder = queryBuilder.Set("quantity", req.Quantity).Where(sq.Eq{"ID": req.BookID})
	}
	if req.Sold != 0 {
		queryBuilder = queryBuilder.Set("sold", req.Sold).Where(sq.Eq{"ID": req.BookID})
	}
	if req.InStock != 0 {
		queryBuilder = queryBuilder.Set("in_stock", req.InStock).Where(sq.Eq{"ID": req.BookID})
	}
	queryBuilder = queryBuilder.Set("status", req.Status).Where(sq.Eq{"ID": req.BookID})

	if err != nil {

		return err
	}
	query, qargs, err := queryBuilder.ToSql()
	if err != nil {
		log.Println(err)
		return err
	}
	row, err := db1.QueryContext(ctx, query, qargs...)
	if err != nil {
		log.Println(err)
		return err
	}

	defer row.Close()
	_, err2 := db1.ExecContext(ctx, query, qargs...)
	if err2 != nil {
		log.Println(err)
		return err
	}
	return nil

}
