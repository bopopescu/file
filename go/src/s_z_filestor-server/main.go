package main

import (
	"fmt"
	"net/http"
	"s_z_filestor-server/db"
	"s_z_filestor-server/db/mysql"
	"s_z_filestor-server/handler"
)

func main() {

	fmt.Println(mysql.Db)


	db.Saveimg()
	//return
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("s_z_filestor-server/static"))))

	http.HandleFunc("/fileupload",handler.Uploadhandler)
	http.HandleFunc("/fileupload/succ",handler.UploadSuccHandler)
	http.HandleFunc("/fileupload/get",handler.GetimgHandler)


	err := http.ListenAndServe(":8657", nil)

	if err!=nil{

		panic(err)
	}
}
