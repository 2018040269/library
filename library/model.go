package library

type Book struct {
	ID     int64  `db:"id"`
	Title  string `db:"title"`
	Number int64  `db:"number"`
}
