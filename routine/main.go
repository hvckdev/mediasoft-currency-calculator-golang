package routine

import (
	"awesomeProject3/table"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
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

			urlRequest := "https://api.exchangerate.host/latest?base=" + table.Currency1

			client := http.Client{}

			request, _ := http.NewRequest("GET", urlRequest, nil)
			resp, _ := client.Do(request)
			var result map[string]interface{}

			json.NewDecoder(resp.Body).Decode(&result)

			newResult := result["rates"].(map[string]interface{})[table.Currency2].(float64)

			db.Exec("UPDATE main SET rate=? WHERE id=?", newResult, table.ID)
		}

		time.Sleep(30 * time.Minute)
	}
}
