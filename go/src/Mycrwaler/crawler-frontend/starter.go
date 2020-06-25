package main

import (

"net/http"

"Mycrwaler/crawler-frontend/controller"

)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("Mycrwaler/crawler-frontend/view")))



	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"Mycrwaler/crawler-frontend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
