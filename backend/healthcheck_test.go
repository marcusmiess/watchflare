package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunHealthcheck(t *testing.T) {
	ok := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	}))
	defer ok.Close()
	if code := runHealthcheck(ok.URL); code != 0 {
		t.Errorf("healthy server: exit code = %d, want 0", code)
	}

	failing := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer failing.Close()
	if code := runHealthcheck(failing.URL); code != 1 {
		t.Errorf("failing server: exit code = %d, want 1", code)
	}

	failing.Close()
	if code := runHealthcheck(failing.URL); code != 1 {
		t.Errorf("unreachable server: exit code = %d, want 1", code)
	}
}
