package mysql

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB


func init(){

	Db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/fileserver?charset=utf8")

	Db.SetMaxOpenConns(100)

	err:=Db.Ping()
	if err!=nil {
		fmt.Print(123123)
		panic(err)
	}



	fmt.Println("数据库连接成功")
}

func DBconn() *sql.DB{
	return Db
}