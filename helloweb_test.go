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
	t.Logf("Special Var: %s", os.Getenv("SPECIAL_VAR"))
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	Hello(w, req)
	t.Logf("Status Code: %d", w.Code)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}

	theBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		log.Fatal(err)
	}

	var HTTPOutput struct {
		Version int `json:"version"`
		Special_var string `json:"special_var"`
	}
	err = json.Unmarshal(theBody, &HTTPOutput)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("Body: %s", theBody)
	t.Logf("Version: %d", HTTPOutput.Version)
	t.Logf("Special Variable: %s", HTTPOutput.Special_var)

	if HTTPOutput.Version != 11 {
		t.Log("Version is not 11, got", HTTPOutput.Version)
		t.Fail()
	}

	var testVar = "ninety-nine"
	if HTTPOutput.Special_var != testVar {
		t.Logf("Special Var is not %s, got %s", testVar, HTTPOutput.Special_var)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	var testVar = "ninety-nine"
	os.Setenv("SECRET_VAR", testVar)

	os.Exit(m.Run())
}
