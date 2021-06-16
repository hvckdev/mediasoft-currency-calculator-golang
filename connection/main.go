package connection

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func Database() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/calculator")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
