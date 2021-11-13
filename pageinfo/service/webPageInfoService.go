package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
)

func Extract(r *domain.Request) (domain.Pageinfo, error) {
	resp, e := scrape(r.URL)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
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
