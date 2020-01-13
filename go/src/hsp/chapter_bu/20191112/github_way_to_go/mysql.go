package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
	//"github.com/go-sql-driver/mysql"
)

const (
	DB_Driver = "root:123456@tcp(192.168.0.110:3306)/consult?charset=utf8"
)

func main() {
	//mysql.Config{}
	opend, db := OpenDB()
	fmt.Println(opend, db)
	if opend {
		fmt.Println("open success")
	} else {
		fmt.Println("open faile:")
	}
	fmt.Println("ok now")
	insertToDB(db)
	QueryFromDB(db)
	return
	DeleteFromDB(db, 10)
	//QueryFromDB(db)
	//DeleteFromDB(db, 1)
	//UpdateDB(db, 5)
	//UpdateUID(db, 5)
	//UpdateTime(db, 4)

}

func DeleteFromDB(db *sql.DB, autid int) {
	stmt, err := db.Prepare("delete from userinfo where autid=?")
	CheckErr(err)
	res, err := stmt.Exec(autid)
	affect, err := res.RowsAffected()
	fmt.Println("删除数据：", affect)
}
func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	db, err := sql.Open("mysql", DB_Driver)
	if err != nil {
		isOpen = false
	} else {
		isOpen = true
	}
	//mysql.Config{}
	CheckErr(err)
	return isOpen, db
}

func insertToDB(db *sql.DB) {
	stmt, err := db.Prepare("insert test1 set name=?")
	CheckErr(err)
	res, err := stmt.Exec("研发中心")
	CheckErr(err)
	id, err := res.LastInsertId()
	CheckErr(err)
	if err != nil {
		fmt.Println("插入数据失败")
	} else {
		fmt.Println("插入数据成功：", id)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	fmt.Println(t)
	return str
}

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetNowtimeMD5() string {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	return GetMD5Hash(timestamp)
}

func QueryFromDB(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM test1")
	CheckErr(err)
	if err != nil {
		fmt.Println("error:", err)
	} else {
	}

	for rows.Next() {
		var id string
		var name string

		CheckErr(err)
		err = rows.Scan(&id, &name)
		fmt.Println(id)
		fmt.Println(name)

		fmt.Println("----------" +
			"++++++")
	}
}
