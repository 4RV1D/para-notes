package main

import (
    "os"
	"fmt"
	"io/ioutil"
	"net/http"
)

func save(w http.ResponseWriter, req *http.Request) {
    if (req.Method == "OPTIONS") {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        return;
    }

    if (req.Method != "POST") {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return;
    }

    body, err := ioutil.ReadAll(req.Body)
    if (err != nil) {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return;
    }

    err = os.WriteFile("data.json", body, 0644);
    if (err != nil) {
        http.Error(w, "Error saving request body", http.StatusInternalServerError)
        return;
    }

    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{ "success": true }`)
}

