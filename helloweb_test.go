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
	req, _ := http.NewRequest("GET", "/version", nil)
	w := httptest.NewRecorder()
	version(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

	result, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var version string = "12"
	stringedResult := strings.TrimSpace(string(result))
	compareOutcome := strings.Compare(version, stringedResult)
	if compareOutcome != 0 {
		t.Log("Result is not 12, got", string(result))
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
