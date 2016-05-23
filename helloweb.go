package main

import (
	"io"
	"net/http"
	"strconv"
)

var VERSION int = 11

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strconv.Itoa(VERSION))
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}
