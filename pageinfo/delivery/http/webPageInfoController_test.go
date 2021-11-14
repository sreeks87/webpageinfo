package delivery_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	delivery "github.com/sreeks87/webpageinfo/pageinfo/delivery/http"
	"github.com/stretchr/testify/assert"
)

func TestHeartBeat(t *testing.T) {
	req, err := http.NewRequest("GET", "/heartbeat", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delivery.Heartbeat)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestWebPageInfoEmptyReq(t *testing.T) {
	r := []byte(`{}`)
	req, err := http.NewRequest("POST", "/webpageinfo", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delivery.Webpageinfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestWebPageInfoEmptyReq2(t *testing.T) {
	req, err := http.NewRequest("POST", "/webpageinfo", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delivery.Webpageinfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestWebPageInfoEmptyReq3(t *testing.T) {
	r := []byte(nil)
	req, err := http.NewRequest("POST", "/webpageinfo", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delivery.Webpageinfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestWebPageInfoEmptyReq4(t *testing.T) {
	r := []byte(`{"url":"http://www.example.com"}`)
	req, err := http.NewRequest("POST", "/webpageinfo", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delivery.Webpageinfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusBadRequest)
}
