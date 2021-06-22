package test

import (
	"awesomeProject3/models"
	"awesomeProject3/pg"
	"context"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestDB(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	err := pg.Connect(cancel)
	if err != nil {
		t.Errorf("error connecting db")
	}

	defer func(DB *sqlx.DB) {
		err := DB.Close()
		if err != nil {
			t.Errorf("error closing db")
		}
	}(pg.DB)

	q := "SELECT * FROM currency"

	var m models.Currency

	err = pg.DB.GetContext(ctx, &m, q)
	if err != nil {
		t.Errorf("error executing query")
	}
}
