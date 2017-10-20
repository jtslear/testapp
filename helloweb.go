package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var VERSION int = 14

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(VERSION))
}

func hello(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, hostname)
}

func main() {
	http.HandleFunc("/version", version)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
