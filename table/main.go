package table

type Main struct {
	ID        int     `db:"id"`
	Currency1 string  `db:"currency1"`
	Currency2 string  `db:"currency2"`
	Rate      float64 `db:"rate"`
	UpdatedAt string  `db:"updated_at"`
}
