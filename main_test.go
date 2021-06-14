package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Homepage)

	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}

func TestLogin(t *testing.T) {

	var payload = []byte(`{"username":"testLogin", "password":"testLogin"}`)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}

func TestContent(t *testing.T) {

	var payload = []byte(`{"token":"2d2f3349-e9a1-4b73-bd18-b236ac1dc02a"}`)

	req, err := http.NewRequest("POST", "/content", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Content)

	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}

func TestHealthcheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Healthcheck)

	handler.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}
}
