package main

import (
	"database/sql"
)

func initDB() (err error) {
	dsn := "root:lwt520...@tcp(127.0.0.1:3306)/lw"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//与数据库连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
