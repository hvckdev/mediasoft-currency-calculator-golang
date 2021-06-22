package main

import (
	"awesomeProject3/consts"
	"awesomeProject3/pg"
	"awesomeProject3/routes"
	"awesomeProject3/routine"
	"context"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	err := pg.Connect(cancel)
	if err != nil {
		panic(err)
	}

	err = pg.Migrate()
	if err != nil {
		panic(err)
	}

	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(pg.DB)

	r := routes.SetupRouter()

	err = r.Run(consts.RConfig.Host + ":" + consts.RConfig.Port)
	if err != nil {
		panic(err)
	}

	go routine.UpdateCurrencies(ctx, consts.UpdateTimeout)
}
