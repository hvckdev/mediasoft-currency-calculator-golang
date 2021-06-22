package pg

import (
	"awesomeProject3/consts"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

func Connect(cancel context.CancelFunc) error {
	db, err := sqlx.Open(consts.DbConfig.Driver, consts.DbSource)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	if db == nil {
		cancel()
		fmt.Println("DB is nil")
	}

	DB = db

	return nil
}

func Migrate() error {
	driver, err := postgres.WithInstance(DB.DB, &postgres.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		consts.DbConfig.DBName, driver)

	if err != nil {
		return err
	}

	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}
