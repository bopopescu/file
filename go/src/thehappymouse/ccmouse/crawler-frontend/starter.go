package main

import (

"net/http"

"thehappymouse/ccmouse/crawler-frontend/controller"

)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("thehappymouse/ccmouse/crawler-frontend/view")))



	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"thehappymouse/ccmouse/crawler-frontend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
