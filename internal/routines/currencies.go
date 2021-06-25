package routines

import (
	"awesomeProject3/internal/models"
	"awesomeProject3/internal/pg"
	"awesomeProject3/pkg/currencies"
	"context"
	"fmt"
	"log"
	"time"
)

func UpdateCurrencies(ctx context.Context, cancel context.CancelFunc, timeout time.Duration) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("ctx done")
		case <-time.After(time.Minute * timeout):
			dn := time.Now()

			rows, err := pg.DB.QueryxContext(ctx, "SELECT * FROM currency")
			if err != nil {
				cancel()
				panic(err)
			}

			for rows.Next() {
				var model models.Currency
				err := rows.StructScan(&model)
				if err != nil {
					panic(err)
				}

				rate, err := currencies.GetRate(model.Currency1, model.Currency2)
				if err != nil {
					panic(err)
				}

				_, err = pg.DB.ExecContext(ctx, "update currency set rate=$1, updated_at=$2 where id=$3", rate, dn, model.ID)
				if err != nil {
					panic(err)
				}
			}

			fmt.Println(dn.String() + " | currency have been updated")
		}
	}
}
