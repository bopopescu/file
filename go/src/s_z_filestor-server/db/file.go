package db

import (
	"s_z_filestor-server/db/mysql"
	"fmt"
	"time"

	_"github.com/go-sql-driver/mysql"

)

func Saveimg() {
	//

	x:=fmt.Sprintf("%x",time.Now().Unix())
	//x:=fmt.Sprintf("%x",time.Now().Unix())
	stmt, e := mysql.DBconn().Prepare(
		"insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`," +
			"`file_addr`) values (?,?,?,?)")

	//result, i := mysql.Db.Exec("select * from tbl_file;")
	//
	//fmt.Println(result,i)

	fmt.Println(e)

	defer stmt.Close()

	stmt.Exec(x,2,3,4)
}


