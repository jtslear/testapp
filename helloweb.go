package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Version indicates the version of this application
var Version = 13

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(Version))
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
