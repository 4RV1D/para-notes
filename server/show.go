package main

import (
	"fmt"
	"net/http"
	"os"
)

func show(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}
