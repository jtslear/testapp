package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	var secretVar = "SECRET_VAR"
	var testVar = "ninety-nine"
	os.Setenv(secretVar, testVar)
}

func TestOutputVersion(t *testing.T) {
	var secretVar = "SECRET_VAR"
	val, ok := os.LookupEnv(secretVar)
	if !ok {
		fmt.Printf("%s is not set\n", secretVar)
	} else {
		fmt.Printf("%s: %s\n", secretVar, val)
	}
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

	if HTTPOutput.Special_var != val {
		t.Logf("Special Var is not %s, got %s", val, HTTPOutput.Special_var)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
