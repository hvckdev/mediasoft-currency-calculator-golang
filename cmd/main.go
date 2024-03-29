package main

import (
	"awesomeProject3/consts"
	"awesomeProject3/internal/pg"
	"awesomeProject3/internal/router"
	"awesomeProject3/internal/routines"
	"context"
	"github.com/jmoiron/sqlx"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	err := pg.Connect(ctx, cancel)
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

	go routines.UpdateCurrencies(ctx, cancel, consts.UpdateTimeout)

	r := router.SetupRouter()

	err = r.Run(consts.RConfig.Host + ":" + consts.RConfig.Port)
	if err != nil {
		panic(err)
	}
}
