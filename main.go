package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	body, err := ioutil.ReadAll(r.Body)
	method := r.Method

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Request received \n method: %s \n path: %s \n body: %s \n", method, path, body)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
