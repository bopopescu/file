package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"log"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "192.168.0.110"
	PORT     = 3306
	DATABASE = "consult"
)

type User struct {
	ID   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	//由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
}

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}

	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(64)                   //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数

	err2 := DB.Ping()

	go func() {

	}()
	fmt.Println("ping err", err2)
	queryMulti(DB)

	queryOne(DB)

}
func queryOne(DB *sql.DB) {
	for i := 1; i < 150; i++ {
		fmt.Println("query times:", i)
		user := new(User)
		row := DB.QueryRow("select id,name from test1 where id=?", i-1)
		//continue
		var s string
		var sd string

		err2 := row.Scan(&s, &sd)
		if err2 == sql.ErrNoRows {
			log.Println("There is not row")
		} else {
			fmt.Println(err2)
		}

		fmt.Println(s, "---", sd, err2, 6666)
		return

		if err := row.Scan(&user.ID, &user.Name); err != nil {
			fmt.Printf("scan failed, err:%v", err)
			//return
		} else {
			fmt.Printf("scan succ, err:%v", err)

		}

		fmt.Println(*user)
	}
}

func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from test1 where id > ?", 3)
	cols, err := rows.Columns()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cols)
	fmt.Printf("cols type:%T\n", cols)

	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))

	testval := 321
	testlove := &testval

	fmt.Println(testlove)
	*testlove =32

	fmt.Println(testval,testlove)


	for i := range vals {
		scans[i] = &vals[i]
	}


	fmt.Println(777, scans ,vals,"\n")

	var results []map[string]string

	for rows.Next() {
		fmt.Println(123, scans, vals)

		err = rows.Scan(scans...)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(213, scans,*scans[0].(*[]byte), vals,"\n")
		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}

	fmt.Println(999, results,"\n")
	for k, v := range results {
		fmt.Println(k, v)
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()

	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	fmt.Println(rows)

	var s []byte
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &s) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}

		fmt.Printf("%p--for %v,", user.Name)
		fmt.Println("SAD", *user, (*user).ID, (*user).Name)
	}
}

func insertData(DB *sql.DB) {
	result, err := DB.Exec("insert INTO users(name,age) values(?,?)", "YDZ", 23)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected() //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}
