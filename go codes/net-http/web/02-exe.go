package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", greetings)
	http.ListenAndServe(":8080", nil)
}

func greetings(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello,%s!", name)

}
