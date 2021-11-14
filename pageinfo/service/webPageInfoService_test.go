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

func TestValidate(t *testing.T) {
	request := domain.Request{
		URL: "http://www.example.com",
	}
	service := service.NewExtractorService(request)

	e := service.Validate()
	assert.Equal(t, e, nil)
}

func TestValidate2(t *testing.T) {
	request := domain.Request{
		URL: "1234566",
	}
	service := service.NewExtractorService(request)

	e := service.Validate()
	assert.NotEqual(t, e, nil)
}

func TestValidate3(t *testing.T) {
	request := domain.Request{
		URL: "",
	}
	service := service.NewExtractorService(request)

	e := service.Validate()
	assert.NotEqual(t, e, nil)
}
