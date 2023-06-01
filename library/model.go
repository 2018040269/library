package library

type Book struct {
	ID    int64   `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
}
