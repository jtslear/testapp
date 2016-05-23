package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestOutputVersion(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	Hello(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

	result, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var eleven string = "11"
	stringedResult := strings.TrimSpace(string(result))
	compareOutcome := strings.Compare(eleven, stringedResult)
	if compareOutcome != 0 {
		t.Log("Result is not 11, got", string(result))
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
