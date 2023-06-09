package library

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:8888)/test_api"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func queryAlllBook() (bookList []*Book, err error) {

	sqlStr := "select id,title ,number from book"

	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Printf("查询信息失败err=%v\n", err)
		return
	}
	return
}

func querySingalBook(id int64) (book Book, err error) {

	sqlstr := "select id,title ,number from book where id=? "
	err = db.Get(&book, sqlstr, id)
	if err != nil {
		fmt.Printf("查询信息失败1111err=%v\n", err)
		return
	}
	return
}

func insertAlllBook(title string, number float64) (err error) {

	sqlStr := "insert into book(title,number) values (?,?)"
	_, err = db.Exec(sqlStr, title, number)
	if err != nil {
		fmt.Printf("插入信息失败err=%v\n", err)
		return
	}
	//var LastInsertId int64
	//LastInsertId, err = ret.LastInsertId() // 新插入数据的id
	//if err != nil {
	//	fmt.Printf("get lastinsert ID failed, err: %v\n", err)
	//	return 0, err
	//}
	return //LastInsertId, nil

}

func deleteBook(id int64) (err error) {

	sqlStr := "delete from book where id=?"

	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("删除信息失败err=%v\n", err)
		return
	}
	return
}

func updateBook(title string, number float64, id int64) (err error) {

	sqlStr := "update book set title=?,number=? where id=?"

	_, err = db.Exec(sqlStr, title, number, id)
	if err != nil {
		fmt.Printf("更新信息失败err=%v\n", err)
		return
	}
	return
}
