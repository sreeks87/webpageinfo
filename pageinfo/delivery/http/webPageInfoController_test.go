package delivery_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	delivery "github.com/sreeks87/webpageinfo/pageinfo/delivery/http"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
	"github.com/sreeks87/webpageinfo/pageinfo/service/mocks"
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
	mockSvc := new(mocks.Service)
	p := domain.Pageinfo{
		HTMLVersion: "4",
		PageTitle:   "title",
		HeadingData: domain.Head{
			H1: 1,
			H2: 2,
			H3: 3,
			H4: 4,
			H5: 5,
			H6: 6,
		},
		LinkData: domain.Links{
			InternalLinks:     1,
			ExternalLinks:     2,
			InaccessibleLinks: 3,
		},
		LoginForm: false,
		Error:     "",
	}
	mockSvc.On("Extract").Return(p, nil)
	r := []byte(`{"url":"http://www.example.com"}`)
	req, err := http.NewRequest("POST", "/webpageinfo", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(delivery.Webpageinfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestWebPageInfoEmptyReq5(t *testing.T) {
	mockSvc := new(mocks.Service)
	p := domain.Pageinfo{}
	mockSvc.On("mockSvc.Extract").Return(p, errors.New("error"))
	r := []byte(`{"url":"1234"}`)
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
