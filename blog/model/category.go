package model

//定义category结构体
//id, category_name, category_no, create_time, update_time

type Category struct {
	CategoryID   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"`
}
