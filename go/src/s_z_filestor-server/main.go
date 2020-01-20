package main

import (
	"net/http"
	"s_z_filestor-server/handler"
)

func main() {


	http.HandleFunc("/fileupload",handler.Uploadhandler)
	http.HandleFunc("/fileupload/succ",handler.UploadSuccHandler)
	http.HandleFunc("/fileupload/get",handler.GetimgHandler)


	err := http.ListenAndServe(":8657", nil)

	if err!=nil{

		panic(err)
	}
}
