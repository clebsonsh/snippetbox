package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clebsonsh/snippetbox/internal/assert"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	got := rs.Header.Get("Content-Security-Policy")
	want := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, got, want)

	got = rs.Header.Get("Referrer-Policy")
	want = "origin-when-cross-origin"
	assert.Equal(t, got, want)

	got = rs.Header.Get("X-Content-Type-Options")
	want = "nosniff"
	assert.Equal(t, got, want)

	got = rs.Header.Get("X-Frame-Options")
	want = "deny"
	assert.Equal(t, got, want)

	got = rs.Header.Get("X-XSS-Protection")
	want = "0"
	assert.Equal(t, got, want)

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(body), "OK")
}
