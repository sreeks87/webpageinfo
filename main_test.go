package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleUI(t *testing.T) {
	// create a dummy http serverfor tests

	r, e := http.NewRequest("GET", "/ui", nil)

	// handle any error from above
	if e != nil {
		t.Fatal(e)
	}

	// create a new recorder to test the service
	rec := httptest.NewRecorder()

	// the actual handler to test
	handler := http.HandlerFunc(ui)
	// serve the http server
	handler.ServeHTTP(rec, r)

	assert.Equal(t, rec.Code, 200)
}

func TestHandleUINotExists(t *testing.T) {
	// create a dummy http serverfor tests

	r, e := http.NewRequest("GET", "/uinonexistent", nil)

	// handle any error from above
	if e != nil {
		t.Fatal(e)
	}

	// create a new recorder to test the service
	rec := httptest.NewRecorder()

	// the actual handler to test
	handler := http.HandlerFunc(ui)
	// serve the http server
	handler.ServeHTTP(rec, r)

	assert.Equal(t, rec.Code, 404)
}
