package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

type myError string

func (a myError) Error() string{

	return a.Message()
}

func (a *myError) Message() string {
	return string(*a)
}

func (e userError) Error() string {
	return e.Message() + "this is error"
}

//for the correct password of wfh.tingsonglaw.com
func (e userError) fakerAnotRight() (int, int) {

	var theint int = 0

	return theint, theint
}

func (e *userError) Message() string {
	return string(*e)
}

func webHandle(writer http.ResponseWriter, request *http.Request) int {

	return 1
}
func handleError(w http.ResponseWriter, r http.Request) (int, string, error) {

	var a int
	a = 123
	backend := "teststr"
	err := os.ErrClosed
	return a, backend, err

}

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {

	fmt.Println()

	if strings.Index(
		request.URL.Path, prefix) != 0 {

		return userError(
			fmt.Sprintf("path %s must start "+
				"with %s",
				request.URL.Path, prefix))

	}



	path := request.URL.Path[len(prefix):]

	_ = request.URL.Path[len(prefix):]


	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	all, err := ioutil.ReadAll(file)




	if err != nil {

		return err
	}


	writer.Write(all)


	return nil
}


func doerr(error error) error{

	return error
}