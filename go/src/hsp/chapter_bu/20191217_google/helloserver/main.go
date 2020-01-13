package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request) {
		fmt.Fprintf(writer,
			"<h1>hello y m t fk  %s </h1>",
			request.FormValue("name"))
		//writer.Write([]byte("fuckyou hello world"))
	})

	http.ListenAndServe(":8888", nil)

}
