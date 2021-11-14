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

func NewExtractorService(r domain.Request) domain.Service {
	return &extractorSVC{
		request: r,
	}
}

func (svc *extractorSVC) Extract() (domain.Pageinfo, error) {

	e := svc.Validate()
	if e != nil {
		log.Println("request validation failed")
		return domain.Pageinfo{}, e

	}
	resp, e := svc.Scrape()
	if e != nil {
		return domain.Pageinfo{}, e
	}
	defer resp.Body.Close()
	body := resp.Body
	if e != nil {
		return domain.Pageinfo{}, nil
	}
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return domain.Pageinfo{}, err
	}
	parser := NewParserSvc(doc)
	result, err := parser.ParsePage(svc.request.URL)
	if err != nil {
		return domain.Pageinfo{}, err
	}
	return result, nil
}

func (svc *extractorSVC) Scrape() (*http.Response, error) {
	log.Println("scraping ", svc.request.URL)
	resp, e := http.Get(svc.request.URL)
	if e != nil {
		return nil, e
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not get response from url : ", svc.request.URL)
	}
	return resp, nil
}

func (svc *extractorSVC) Validate() error {
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
