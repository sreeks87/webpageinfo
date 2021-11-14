package service_test

import (
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
	"github.com/sreeks87/webpageinfo/pageinfo/service"
	"github.com/sreeks87/webpageinfo/pageinfo/service/mocks"
	"github.com/stretchr/testify/assert"
)

// function to test the  Exctract method
// pass a real body, expects no error in response
func TestExtract(t *testing.T) {
	request := domain.Request{
		URL: "http://www.example.com",
	}
	service := service.NewExtractorService(request)
	mockSvc := new(mocks.Service)

	mockSvc.On("Validate").Return(nil)
	mockSvc.On("Scrape").Return(nil, nil)
	_, e := service.Extract()
	assert.Equal(t, e, nil)
}

// function to test Extract with invalid url
// expected error
func TestExtractError(t *testing.T) {
	request := domain.Request{
		URL: "http",
	}
	service := service.NewExtractorService(request)
	mockSvc := new(mocks.Service)

	mockSvc.On("Validate").Return(errors.New("invalid"))
	mockSvc.On("Scrape").Return(nil, nil)
	_, e := service.Extract()
	assert.NotEqual(t, e, nil)
}

// funtion to test the Extract method,
// when the internal call to Scrape returns error.
func TestExtractErrorScrape(t *testing.T) {
	request := domain.Request{
		URL: "http://www.exampleexapleexampleexample.com",
	}
	service := service.NewExtractorService(request)
	mockSvc := new(mocks.Service)

	mockSvc.On("Validate").Return(nil)
	mockSvc.On("Scrape").Return(nil, errors.New("error"))
	_, e := service.Extract()
	assert.NotEqual(t, e, nil)
}

// function to test Scrape method
// mock the external http call to rstrict network call
// dummy http service returns bad request
// expects not nil error
func TestScrape(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "http://www.exampleexapleexampleexample.com",
		httpmock.NewStringResponder(400, "response"))
	defer httpmock.DeactivateAndReset()
	request := domain.Request{
		URL: "http://www.exampleexapleexampleexample.com",
	}
	service := service.NewExtractorService(request)

	_, e := service.Scrape()
	assert.NotEqual(t, e, nil)
}

// function to test Scrape method
// provide valid inputs and mock the external http call
// expects nil error

func TestScrape2(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "http://www.example.com",
		httpmock.NewStringResponder(200, "response"))
	defer httpmock.DeactivateAndReset()
	request := domain.Request{
		URL: "http://www.example.com",
	}
	service := service.NewExtractorService(request)

	_, e := service.Scrape()
	assert.Equal(t, e, nil)
}

// function to test Validate method
// provide valid inputs
// expects nil error
func TestValidate(t *testing.T) {
	request := domain.Request{
		URL: "http://www.example.com",
	}
	service := service.NewExtractorService(request)

	e := service.Validate()
	assert.Equal(t, e, nil)
}

// function to test Validate method
// provide invalid inputs
// expects not nil error

func TestValidate2(t *testing.T) {
	request := domain.Request{
		URL: "1234566",
	}
	service := service.NewExtractorService(request)

	e := service.Validate()
	assert.NotEqual(t, e, nil)
}

// function to test Validate method
// provide invalid inputs with empty url
// expects not nil error
func TestValidate3(t *testing.T) {
	request := domain.Request{
		URL: "",
	}
	service := service.NewExtractorService(request)

	e := service.Validate()
	assert.NotEqual(t, e, nil)
}
