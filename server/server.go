package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/show", show)
	http.HandleFunc("/save", save)

	http.ListenAndServe(":9000", nil)
}
