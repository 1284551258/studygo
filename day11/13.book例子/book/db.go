package book

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func IniDB() (err error) {
	addr := "root:root@tcp(11.2.2.128:3306)/go_learn"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		fmt.Println("Open mysql failed，err:", err)
		return
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return
}

func QueryAllBook() (bookList *[]Book, err error) {
	bookList = new([]Book)
	sqlStr := "SELECT id, title, price FROM book;"
	err = db.Select(bookList, sqlStr)
	if err != nil {
		fmt.Println("query all book failed,err:", err)
		return nil, err
	}
	return
}

func InsertBook(title string, price int64) (err error) {
	sqlStr := "INSERT INTO book (title, price) VALUES(?, ?);"
	res, err := db.Exec(sqlStr, title, price)
	if err != nil {
		fmt.Println("insert failed,err:", err)
		return err
	}
	fmt.Println("insert successful,id:", res)
	return
}
func DeleteBook(id int64) error {
	sqlStr := "DELETE FROM book WHERE id=?;"
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed,err:", err)
		return err
	}
	fmt.Println("delete successful,id:", res)
	return err
}
