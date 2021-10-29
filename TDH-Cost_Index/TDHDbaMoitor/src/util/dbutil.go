package util

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func GetdbConn (mysqlconn string) *sql.DB{
	db, err := sql.Open("mysql", mysqlconn+"/workflow_workflow1?charset=utf8")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
