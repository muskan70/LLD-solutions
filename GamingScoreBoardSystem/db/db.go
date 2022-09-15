package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySqlDBConnection() (*sql.DB, error) {
	str := "root:1234@13@tcp(127.0.0.1:3306)/intuitMC"
	conn, err := sql.Open("mysql", str)
	if err != nil {
		log.Println("Not able to connect", err.Error())
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		log.Println("Not able to connect", err.Error())
		return nil, err
	}
	return conn, nil
}
