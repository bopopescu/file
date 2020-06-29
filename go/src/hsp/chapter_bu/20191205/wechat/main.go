package main

import (
	"fmt"
	"hsp/chapter_bu/20191205/wechat/wx"
	"log"
	"net/http"
	"time"
)

const (
	logLevel = "dev"
	port     = 8081
	token    = "kosfe9rtue9r1"
)

func get(w http.ResponseWriter, r *http.Request) {

	client, err := wx.NewClient(r, w, token)

	if err != nil {
		log.Println(err)
		w.WriteHeader(405)
		return
	}

	if len(client.Query.Echostr) > 0 {
		w.Write([]byte(client.Query.Echostr))
		return
	}

	w.WriteHeader(406)
	return
}

func post(w http.ResponseWriter, r *http.Request) {

	client, err := wx.NewClient(r, w, token)

	if err != nil {
		log.Println(err)
		w.WriteHeader(403)
		return
	}

	client.Run()
	return
}

func handlerjob(w http.ResponseWriter,r *http.Request){

}

type Apiserver struct {
	httpServer *http.Server
}

var (
	G_Apiserver *Apiserver
)
func main() {
	mux :=http.NewServeMux()  //这个继承里servehttp 所以也可以当成handler
	mux.HandleFunc("/fuck",handlerjob)


	server := http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        &httpHandler{},
		//Handler:       mux ,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}


	G_Apiserver = &Apiserver{
		httpServer:&server,
	}

	fmt.Println(G_Apiserver)


	log.Println(fmt.Sprintf("Listen: %d", port))

	//log fatal means exit program
	log.Fatal(server.ListenAndServe())
}
