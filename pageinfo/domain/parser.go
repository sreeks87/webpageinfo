package domain

// this is the document parser interface
// it is responsible for the parsing and calculation of
// title/html version/headings/links/title
type Parser interface {
	ParsePage(string) (Pageinfo, error)
	ParseHtmlVersion() (string, error)
	ParseHead() (Head, error)
	ParseLinks(string) (Links, error)
	ParseTitle() (string, error)
}
