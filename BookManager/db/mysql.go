package db

import (
	"context"
	"database/sql"
	"dotpe/demo/utils"
	"fmt"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var Client *sql.DB

func Init() (*sql.DB, error) {
	str := "root:Ayushdd@123@tcp(127.0.0.1:3306)/library"
	var err error
	db, err := sql.Open("mysql", str)
	if err == nil {
		return db, err
	}
	if err != nil {
		fmt.Println("Not able to connect")
	}
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Not able to connect")
		return db, pingErr
	}

	return db, err
}
func WriteDBError(ctx context.Context, err error, db, etype string) {
	_, file, line, _ := runtime.Caller(1)
	utils.DLogger.WriteLogs(ctx, logrus.Fields{
		"caller": fmt.Sprintf("%s:%s:%s:%d", db, etype, file, line),
		"error":  err.Error(),
	}, logrus.ErrorLevel, "MySQLError")
}
