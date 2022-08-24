package main

//导入mysql驱动
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() (err error) {
	dsn := "root:root@tcp(11.2.2.128:3306)/sql_test"
	//获得数据库连接
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("数据库连接信息不合法，%v\n", err)
		return
	}
	//尝试连接数据库
	if err = db.Ping(); err != nil {
		fmt.Printf("数据库连接信息不正确，%v\n", err)
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// queryOne ...
func queryOne(id int) {

	var u user
	sqlStr := `select * from user where id = ?`

	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%#v", u)

}
func queryMore(id int) {
	sqlStr := `select * from user where id > ?`
	r, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query faild，err:%v", err)
		return
	}
	defer r.Close() //注意释放连接
	for r.Next() {
		var u user
		err := r.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("Scan faild,err:%v", err)
			return
		}
		fmt.Printf("%#v\n", u)
	}
}

func insertOne(u user) {
	sqlStr := `INSERT INTO user(name, age) VALUES (?,?)`
	r, err := db.Exec(sqlStr, u.name, u.age)
	if err != nil {
		fmt.Printf("INSERT Failed,%v", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Printf("Get LastInsertId Failed,%v ", err)
		return
	}
	fmt.Printf("插入id为:%d", id)

}
func updateRow(id int, newName string, newAge int) {

	sqlStr := `UPDATE user
	SET name=?, age=?
	WHERE id=?;`
	r, err := db.Exec(sqlStr, newName, newAge, id)
	if err != nil {
		fmt.Printf("UPDATE Failed,%v", err)
		return
	}
	i, err := r.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected Failed,%v ", err)
		return
	}
	fmt.Printf("受影响的行数为:%d", i)

}
func deleteOne(id int) {
	sqlStr := `DELETE FROM user
	WHERE id=?;`

	r, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("DELETE Failed,%v", err)
		return
	}
	i, err := r.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected Failed,%v ", err)
		return
	}
	fmt.Printf("受影响的行数为:%d", i)
}

func insertPrepare() {
	sqlStr := `INSERT INTO user(name, age) VALUES (?,?)`
	s, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare Failed,err:%v", err)
		return
	}
	defer s.Close()
	u := map[string]int{
		"王五": 50,
		"赵六": 60,
	}

	for name, age := range u {

		r, err := s.Exec(name, age)
		if err != nil {
			fmt.Printf("INSERT Failed,err:%v", err)
			return
		}
		id, err := r.LastInsertId()
		if err != nil {
			fmt.Printf("Get LastInsertId Failed,err:%v", err)
			return
		}
		fmt.Printf("插入id为:%d\n", id)
	}
}

func transAction() {
	sqlStr1 := `UPDATE user
	SET age=age+10
	WHERE id=2;`
	sqlStr2 := `UPDATE user2
	SET age=age-10
	WHERE id=3;`

	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("Begin Failed!,err:%v", err)
		return
	}
	r1, err := tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("执行strsql1报错,err:%v", err)
		return
	}
	i, err := r1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("执行strsql1报错,err:%v", err)
		return
	}
	fmt.Printf("strsql1受影响行数为:%d", i)
	r2, err := tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("执行strsql2报错,err:%v", err)
		return
	}
	i2, err := r2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("执行strsql2报错,err:%v", err)
		return
	}
	fmt.Printf("strsql1受影响行数为:%d", i2)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("提交报错,err:%v", err)
		return
	}
}
func main2() {
	//数据库连接信息
	err := initDb()
	if err != nil {
		fmt.Printf("数据库连接失败，%v\n", err)
		return
	}
	fmt.Println("数据库连接成功！")
	// queryOne(1)
	// queryMore(0)
	// u := user{
	// 	name: "李四",
	// 	age: 40,
	// }
	// insertOne(u)
	// updateRow(1,"zhoulin",9000)
	// deleteOne(1)
	// insertPrepare()
	transAction()

}
