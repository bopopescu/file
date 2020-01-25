package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"hsp/chapter_bu/20191217_google/u2pppw/errhandling/filelistingserver/filelisting"
	//"github.com/sirupsen/logrus"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) http.HandlerFunc {
//func errWrapper(handler appHandler) func(http.ResponseWriter,
//	*http.Request) {

	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic  保护程序panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)

				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			//todo because warn comes from gopm
			//why  cant  find log.warn  method ?
			//is the  package dose not import as the golang/x/tools?
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error()+"22222")

			// user error
			fmt.Println(123)
			if userErr, ok := err.(userError); ok {
				fmt.Println(ok)

				http.Error(writer,
					//"message=>"+userErr.Message()+
					"error=>"+userErr.Error(),
					http.StatusBadRequest)
				return
			}

			fmt.Println(321)

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	//Message() string
}

func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	testFuckHandleWithhttp()

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

func testFuckHandleWithhttp() {
	http.HandleFunc("/fuck123", func(writer http.ResponseWriter,
		request *http.Request) {
		fmt.Fprintf(writer, "Hellofuck ohshit,"+request.URL.Path[1:])

		http.Error(writer, "er", http.StatusInternalServerError)

	})
}
