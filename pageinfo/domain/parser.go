package domain

type Parser interface {
	ParsePage(string) (Pageinfo, error)
	ParseHtmlVersion() (string, error)
	ParseHead() (Head, error)
	ParseLinks(string) (Links, error)
	ParseTitle() (string, error)
}
