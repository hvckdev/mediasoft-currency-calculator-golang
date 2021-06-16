package routine

import (
	"awesomeProject3/currencies"
	"awesomeProject3/table"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func UpdateCurrencies() {
	for {
		db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/calculator")
		if err != nil {
			log.Fatalln(err)
		}

		rows, err := db.Queryx("SELECT * FROM main")
		if err != nil {
			log.Fatalln(err)
		}

		table := table.Main{}

		for rows.Next() {
			err := rows.StructScan(&table)
			if err != nil {
				log.Fatalln(err)
			}

			db.Exec("UPDATE main SET rate=? WHERE id=?", currencies.GetRate(table.Currency1, table.Currency2), table.ID)
		}

		time.Sleep(30 * time.Minute)
	}
}
