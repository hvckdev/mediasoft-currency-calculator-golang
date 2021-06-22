package routine

import (
	"awesomeProject3/currencies"
	"awesomeProject3/models"
	"awesomeProject3/pg"
	"context"
	"log"
	"time"
)

func UpdateCurrencies(ctx context.Context, timeout time.Duration) {
	var model models.Currency
	rate, err := currencies.GetRate(model.Currency1, model.Currency2)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		rows, err := pg.DB.QueryxContext(ctx, "SELECT * FROM currency")
		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {
			err := rows.StructScan(&model)
			if err != nil {
				log.Fatalln(err)
			}

			_, err = pg.DB.ExecContext(ctx, "UPDATE currency SET rate=$1 WHERE id=$2", rate, model.ID)
			if err != nil {
				log.Fatalln(err)
			}
		}

		time.Sleep(timeout * time.Minute)
	}
}
