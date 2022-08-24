package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db2 *sqlx.DB

func initDbx() (err error) {
	dsn := "root:root@tcp(11.2.2.128:3306)/sql_test"
	//获得数据库连接
	db2, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接信息不合法，%v\n", err)
		return
	}
	db2.SetMaxOpenConns(10)
	db2.SetMaxIdleConns(5)
	return
}

type user2 struct {
	ID   int
	Name string
	Age  int
}

func main4() {
	err := initDbx()
	if err != nil {
		fmt.Printf("initDbx failed!,err:%v", err)
		return
	}
	fmt.Println("数据库连接成功！")
	var u2 user2
	sqlStr1 := `SELECT id, name, age
	FROM user where id = 2
	`
	err = db2.Get(&u2, sqlStr1)
	if err != nil {
		fmt.Printf("Get failed!,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", u2)

	var u2List []user2
	sqlStr2 := `SELECT id, name, age
	FROM user 
	`
	err = db2.Select(&u2List, sqlStr2)
	if err != nil {
		fmt.Printf("SELECT failed!,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", u2List)

}
