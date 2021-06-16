package routine

import (
	"awesomeProject3/connection"
	"awesomeProject3/currencies"
	"awesomeProject3/table"
	"log"
	"time"
)

func UpdateCurrencies() {
	for {
		rows, err := connection.Database().Queryx("SELECT * FROM main")
		if err != nil {
			log.Fatalln(err)
		}

		table := table.Main{}

		for rows.Next() {
			err := rows.StructScan(&table)
			if err != nil {
				log.Fatalln(err)
			}

			connection.Database().Exec("UPDATE main SET rate=? WHERE id=?", currencies.GetRate(table.Currency1, table.Currency2), table.ID)
		}

		time.Sleep(30 * time.Minute)
	}
}
