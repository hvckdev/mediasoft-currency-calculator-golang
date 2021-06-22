package consts

import (
	"fmt"
	"time"
)

type BD struct {
	Driver   string
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

type Router struct {
	Host string
	Port string
}

type CurrencyUpdate struct {
	Timeout int
}

var DbConfig = BD{
	Driver:   "postgres",
	Host:     "localhost",
	Port:     3306,
	User:     "hvck",
	Password: "root",
	DBName:   "calculator",
}

var DbSource = fmt.Sprintf(
	"user=%s dbname=%s sslmode=disable",
	DbConfig.User,
	DbConfig.DBName,
)

var RConfig = Router{
	Host: "localhost",
	Port: "3333",
}

const UpdateTimeout time.Duration = 30
