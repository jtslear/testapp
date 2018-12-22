package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	hello(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

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

	sr := strings.TrimSpace(string(result))
	assert.Equal(
		t,
		sr,
		"13",
		"Expected %s, got %s\n", sr, "13",
	)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
