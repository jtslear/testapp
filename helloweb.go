package main

import (
	"encoding/json"
	"net/http"
	"os"
)

var VERSION int = 11

func Hello(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		version int
		special_var string
	}
	var special_var = os.Getenv("SPECIAL_VAR")
	u := Response{version: VERSION, special_var: special_var}
	json.NewEncoder(w).Encode(u)
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}
