package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@/bbs")
	if err != nil {
		log.Fatal(err)
	}
	return
}
