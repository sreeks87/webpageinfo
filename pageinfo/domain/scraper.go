package domain

import "io"

type Scraper interface {
	Scrape(url string) (io.ReadCloser, error)
}
