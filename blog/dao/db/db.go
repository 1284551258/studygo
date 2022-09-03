package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

//Init db初始化函数
func Init(dns string) (err error) {
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}
	//判断是否连接成功
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	return

}
