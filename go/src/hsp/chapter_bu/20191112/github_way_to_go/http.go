package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
func HelloServer2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("fuck you")
	fmt.Fprintf(w, "Hellofuck ohshit,"+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/fuck/", HelloServer2)
	err := http.ListenAndServe("localhost:8080", nil)
	fmt.Println("sssssssss")
	if err != nil {

		log.Fatal("ListenAndServe: ", err.Error())
	}
}
