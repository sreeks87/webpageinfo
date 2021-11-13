package service

import (
	"fmt"
	"net/http"

	"github.com/sreeks87/webpageinfo/pageinfo/domain"
)

func Extract(r domain.Request) (domain.Pageinfo, error) {
	_, e := scrape(r.URL)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	return domain.Pageinfo{}, nil
}

func scrape(url string) (*http.Response, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not get response from url : ", url)
	}
	defer resp.Body.Close()

	return resp, nil
}
