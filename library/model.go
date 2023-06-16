package library

type Book struct {
	ID     int64   `db:"id"` //和数据库联系加一个db的tag
	Title  string  `db:"title"`
	Number float64 `db:"number"`
}
