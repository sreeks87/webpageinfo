package domain

import "net/http"

// Pageinfo is the response struct
// this will contain all the data required by the client
type Pageinfo struct {
	HTMLVersion string `json:"htmlversion"`
	PageTitle   string `json:"pagetitle"`
	HeadingData Head   `json:"headings"`
	LinkData    Links  `json:"links"`
	LoginForm   bool   `json:"loginform"`
	Error       string `json:"error"`
}

// Links is the count of different links
// internal liks count
// external linkscount
// broken links count
type Links struct {
	InternalLinks     int `json:"internallinks"`
	ExternalLinks     int `json:"externallinks"`
	InaccessibleLinks int `json:"inaccessiblelinks"`
}

// Head struct holds the heading count for header h1-h6
type Head struct {
	H1 int `json:"h1count"`
	H2 int `json:"h2count"`
	H3 int `json:"h3count"`
	H4 int `json:"h4count"`
	H5 int `json:"h5count"`
	H6 int `json:"h6count"`
}

// Request is the request body from the client
type Request struct {
	URL string `json:"url"`
}

// Service is the actual service interface and will be responsible for the
// Extraction of data, Scraping/Accessing the
// links and validating the request
type Service interface {
	Extract() (*Pageinfo, error)
	Scrape() (*http.Response, error)
	Validate() error
}
