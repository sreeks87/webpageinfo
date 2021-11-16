package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
)

type extractorSVC struct {
	request domain.Request
}

// NewExtractorService creates and returns a new Extractor service object
func NewExtractorService(r domain.Request) domain.Service {
	return &extractorSVC{
		request: r,
	}
}

// Extract method extracts the various details needed for response
// takes a service object with URL
// input : none
// output :PAgeinfo,error
func (svc *extractorSVC) Extract() (*domain.Pageinfo, error) {
	// call validate tovalidate the request
	e := svc.Validate()
	if e != nil {
		log.Println("request validation failed")
		return nil, e

	}
	// scrape the URL passed in the request object
	resp, e := svc.Scrape()
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	body := resp.Body
	if e != nil {
		return nil, nil
	}
	// create a new goquery document for easy parsing
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	// create anew parserobject to call the different parse methods
	// ParsePage will internally call other required methods
	parser := NewParserSvc(doc)
	result, err := parser.ParsePage(svc.request.URL)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Scrape function scrapes/accesses the URL
// input:  request object and scrapes the response html
// output : returns the http Response object
func (svc *extractorSVC) Scrape() (*http.Response, error) {
	log.Println("scraping ", svc.request.URL)
	resp, e := http.Get(svc.request.URL)
	if e != nil {
		return nil, e
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not get response from url : %s", svc.request.URL)
	}
	return resp, nil
}

// Validate function validates the URL
// takes a string url and responds with error, if invalid
func (svc *extractorSVC) Validate() error {
	// empty url check
	if svc.request.URL == "" {
		return errors.New("no url specified")
	}
	// check if URL is valid
	_, e := url.ParseRequestURI(svc.request.URL)
	if e != nil {
		return errors.New("invalid url")
	}

	return nil

}
