package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

func Init() (*sql.DB, error) {
	str := "root:1234@123@tcp(127.0.0.1:3306)/library"
	var err error
	db, err := sql.Open("mysql", str)
	if err == nil {
		return db, err
	}
	if err != nil {
		fmt.Println("Not able to connect", err.Error())
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Not able to connect", err.Error())
		return db, pingErr
	}

	return db, err
}
