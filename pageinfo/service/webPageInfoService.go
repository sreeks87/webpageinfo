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

func Extract(r *domain.Request) (domain.Pageinfo, error) {

	e := Validate(r)
	if e != nil {
		log.Println("request validation failed")
		return domain.Pageinfo{}, e

	}
	resp, e := scrape(r.URL)
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
	result, err := parsePage(doc, r.URL)
	if err != nil {
		return domain.Pageinfo{}, err
	}
	return result, nil
}

func scrape(url string) (*http.Response, error) {
	log.Println("scraping ", url)
	resp, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not get response from url : ", url)
	}
	return resp, nil
}

func Validate(r *domain.Request) error {
	if r.URL == "" {
		return errors.New("no url specified")
	}
	// check if URL is valid
	_, e := url.ParseRequestURI(r.URL)
	if e != nil {
		return errors.New("invalid url")
	}

	return nil

}
