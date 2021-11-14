package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	r := newRouter()
	mockSer := httptest.NewServer(r)

	resp, e := http.Get(mockSer.URL + "/heartbeat")
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestRouterInvalidRoute(t *testing.T) {
	r := newRouter()
	mockSer := httptest.NewServer(r)

	resp, e := http.Get(mockSer.URL + "/heartbeatdffd")
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, resp.StatusCode, http.StatusNotFound)
}

func TestRouterInvalidMethod(t *testing.T) {
	r := newRouter()
	mockSer := httptest.NewServer(r)

	resp, e := http.Post(mockSer.URL+"/heartbeat", "", nil)
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, resp.StatusCode, http.StatusMethodNotAllowed)
}

func TestRouterInvalidMethod2(t *testing.T) {
	r := newRouter()
	mockSer := httptest.NewServer(r)

	resp, e := http.Get(mockSer.URL + "/webpageinfo")
	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, resp.StatusCode, http.StatusMethodNotAllowed)
}
