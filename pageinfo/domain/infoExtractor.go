package domain

type Pageinfo struct {
	HTMLVersion string `json:"htmlversion"`
	PageTitle   string `json:"pagetitle"`
	HeadingData Head   `json:"headings"`
	LinkData    Links  `json:"links"`
	LoginForm   bool   `json:"loginform"`
	Error       error  `json:"error"`
}

type Links struct {
	InternalLinks     int `json:"internallinks"`
	ExternalLinks     int `json:"externallinks"`
	InaccessibleLinks int `json:"inaccessiblelinks"`
}

type Head struct {
	H1 int `json:"h1count"`
	H2 int `json:"h2count"`
	H3 int `json:"h3count"`
	H4 int `json:"h4count"`
	H5 int `json:"h5count"`
	H6 int `json:"h6count"`
}

type Request struct {
	URL string `json:"url"`
}

// type InfoExtractor interface {
// 	Extract(string) (Pageinfo, error)
// }
