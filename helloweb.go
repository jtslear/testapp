package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	read, err := ioutil.ReadFile("VERSION")
	if err != nil {
		fmt.Println("Error:", err)
	}
	io.WriteString(w, string(read[:]))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
