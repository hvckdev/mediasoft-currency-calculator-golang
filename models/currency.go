package models

type Currency struct {
	ID        int     `db:"id"`
	Currency1 string  `db:"currency1"`
	Currency2 string  `db:"currency2"`
	Rate      float64 `db:"rate"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}
