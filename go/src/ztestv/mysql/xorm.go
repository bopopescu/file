package main

import (
	"fmt"
	"time"


	_"github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

//定义的自定义函数
func ShowTimestamp(now time.Time) string {
	return fmt.Sprintf("%d", now.Unix())
}

func main() {


	var sd = httpHandler{}
fmt.Println(sd)

	Draw()

	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "Douyutv666!@#", "140.143.32.113", "3306", "consult"))
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	fmt.Println(123)

	err = engine.Ping()
	if err != nil {
		panic(err)
	}
	sql_1_1 := "select * from z_consult where id<50"
	results, err := engine.QueryString(sql_1_1)
	fmt.Println(4442,results,err)



	engine.SetMaxIdleConns(2)
	engine.SetMaxOpenConns(222)
	for _,v:=range results{
		//fmt.Println(k,v)

		for kk,vv :=range v{
			fmt.Println(kk,66666      ,vv)

		}
	}


	//engine.ShowExecTime(true)
	//engine.ShowSQL(true)
	//tpl := xorm.Default("./templates", ".tpl")
	//tpl.SetFuncs("test.tpl", xorm.FuncMap{"ShowTimestamp": ShowTimestamp})
	//err = engine.RegisterSqlTemplate(tpl)
	//if err != nil {
	//	panic(err)
	//}
	//result := engine.SqlTemplateClient("test.tpl",&map[string]interface{}{"now":time.Now()}).Query()
}

//模板中编写的sql内容如下
//select * from user{{ShowTimestamp .now}}