package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestOutputVersion(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	Hello(w, req)
	t.Log("Status Code: %s", w.Code)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

	_, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var HTTPOutput struct {
		Version int `json:"version"`
		Special_var string `json:"special_var"`
	}
	json.NewDecoder(w.Body).Decode(&HTTPOutput)
	t.Log("Body: %s", w.Body)
	t.Log("Version: %d", HTTPOutput.Version)
	t.Log("Special Variable: %s", HTTPOutput.Special_var)

	if HTTPOutput.Version != 11 {
		t.Log("Version is not 11, got", HTTPOutput.Version)
		t.Fail()
	}

	var test_var = "ninety-nine"
	os.Setenv("SECRET_VAR", test_var)
	if HTTPOutput.Special_var != test_var {
		t.Logf("Special Var is not %s, got %s", test_var, HTTPOutput.Special_var)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
