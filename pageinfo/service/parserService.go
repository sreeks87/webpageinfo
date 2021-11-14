package service

import (
	"log"
	u "net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
)

// constant to match with the "password" input type
const PASSWORD = "password"

// the struct that implements the parser service
type parserSVC struct {
	doc *goquery.Document
}

// the below function creates and returns a parser object
func NewParserSvc(d *goquery.Document) domain.Parser {
	return &parserSVC{
		doc: d,
	}
}

// ParsePage function to parse the page and find all the details required for the client
// input : url, document via the reciever type
func (parser *parserSVC) ParsePage(url string) (domain.Pageinfo, error) {

	// parse the document and find the headings count
	head, e := parser.ParseHead()
	if e != nil {
		return domain.Pageinfo{}, e
	}
	// parse the document to find the different types of the links
	// external,internal,inaccessible
	links, e := parser.ParseLinks(url)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	// parse the document and find the title
	title, e := parser.ParseTitle()
	if e != nil {
		return domain.Pageinfo{}, e
	}
	// parse the document and find the html version
	version, e := parser.ParseHtmlVersion()
	if e != nil {
		return domain.Pageinfo{}, e
	}
	// parse the document and find the presence of a login form
	login, e := parser.ParseLoginForm()
	if e != nil {
		return domain.Pageinfo{}, e
	}
	// final repsonse based on all the funtions above.
	// this will be the final response
	pageinfo := domain.Pageinfo{
		HTMLVersion: version,
		PageTitle:   title,
		HeadingData: head,
		LinkData:    links,
		LoginForm:   login,
		Error:       "",
	}

	return pageinfo, nil
}

// the ParseHtmlVersion functions extracts the version number of the document
// assumption : the root html node will have the doctype in one of the belwo formats
// <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
// <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN""http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
// the common part of these version strings is as below
// common part --> -//W3C//DTD [----]//EN
// the below function extracts this string and returns as the version
// Ref : https://www.w3.org/QA/2002/04/valid-dtd-list.html
func (parser *parserSVC) ParseHtmlVersion() (string, error) {
	commonStart := "-//W3C//DTD"
	commonEnd := "//EN"
	// default version is set as 5
	version := "HTML 5.0"
	// checking if the document nodes contains one or more children
	// if yes, find if the child has non zero attributes count
	// assumption: t the 0th attribute of the 0th child will be the doctype
	// replace the common parts of the string to extractonly the version string
	if len(parser.doc.Selection.Nodes) > 0 {
		if len(parser.doc.Selection.Nodes[0].FirstChild.Attr) > 0 {
			vString := parser.doc.Selection.Nodes[0].FirstChild.Attr[0].Val
			version = strings.Replace(vString, commonStart, "", 1)
			version = strings.Replace(version, commonEnd, "", 1)
		}
	}
	return version, nil
}

// the ParseLoginForm function checks the presense of a login form in the html document
// assumption : the login form will contain a password type of input
// if yes, then login form present == true, else false
// breaks with the first occurence of the password type
func (parser *parserSVC) ParseLoginForm() (bool, error) {
	found := false
	parser.doc.Find("input").EachWithBreak(func(i int, s *goquery.Selection) bool {
		in, _ := s.Attr("type")
		log.Println("form input type ", in)
		if v, _ := s.Attr("type"); v == PASSWORD {
			found = true
			return false
		}
		return true
	})
	log.Println("password ", found)
	return found, nil
}

// the ParseHead function parses the document and gets the count of the headings from h1-h6
// input : html document
// output: the Head struct with the count
func (parser *parserSVC) ParseHead() (domain.Head, error) {
	head := map[string]int{"h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
	for k, _ := range head {
		log.Println("-------------------------HEADING ", k, "---------------------")
		parser.doc.Find(k).Each(func(i int, s *goquery.Selection) {
			head[k] += 1
		})
	}
	headCount := domain.Head{
		H1: head["h1"],
		H2: head["h2"],
		H3: head["h3"],
		H4: head["h4"],
		H5: head["h5"],
		H6: head["h6"],
	}

	return headCount, nil
}

// The ParseLinks function does 3 things
// 1. it checks if the link is internal
// assumption : a link is considered internal if it starts in the
// href with / or the same basepath of url
// 2. if the link is external
// assumption : a link is external if the base path do not match and
// has http as the prefix, this was done because some html pages were seen with comment ids as liks
// and these were getting treated as external liks
// 3. if the link is accessible
// make a get call to the link concurrenlty.
// the input to this concurrent function will be the links map urlSet
// this is to make sure the links are not duplicated in the input to the function
func (parser *parserSVC) ParseLinks(url string) (domain.Links, error) {
	log.Println("-------------------------Links---------------------")
	parsedURL, _ := u.Parse(url)
	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	var exLink int
	var inLink int
	// the map to store the external and internal links
	// using a map over list to remove duplication of links
	urlSet := make(map[string]bool)
	var inAccessLink = make(chan int)
	inAccCount := 0
	// finding all anchor tags
	parser.doc.Find("a").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Attr("href")
		// assumption made : an internal link contains either the url as the starting part or /
		// https://www.google.com/contact and /blog will be considered as internal link for url ->www.google.com
		// if the identified link has the same basepath then it is internal link
		// if thelink has a prefix / then assuming it to be an internal link without basepath
		// anything else is an external link
		if strings.HasPrefix(t, baseURL) {
			log.Println("internal link ->", t)
			inLink += 1
			if !urlSet[t] {
				urlSet[t] = true
			}
		} else if strings.HasPrefix(t, "/") {
			log.Println("internal link without url->", t)
			inLink += 1
			if !urlSet[t] {
				urlSet[baseURL+t] = true
			}
		} else {
			if strings.HasPrefix(t, "http") {
				log.Println("external link ->", t)
				exLink += 1
				if !urlSet[t] {
					urlSet[t] = true
				}
			}
		}
	})

	log.Println("internal links count ", inLink)
	log.Println("external links count ", exLink)
	log.Println("links to access ", urlSet)
	// for each elemnt in the urlSet, concurrenlty check the accessibility of the link
	// if inaccessible, then increment the inaccessible counter by 1
	var wg sync.WaitGroup
	for link, _ := range urlSet {
		log.Println("tring to access ", link)
		wg.Add(1)
		go concurrentLinkAccess(link, &wg, inAccessLink)
	}

	// a simple channel monitor to monitor the channel and close when appropriate to do so
	go checkChan(&wg, inAccessLink)

	// each concurrent function will return a 0 or 1, based on the accessibility
	// we keep counting and the final value will become the total inaccesible link count
	for count := range inAccessLink {
		inAccCount += count
		log.Println("Count --->", count)
		log.Println("InAcc count --->", inAccCount)
	}
	links := domain.Links{
		InternalLinks:     inLink,
		ExternalLinks:     exLink,
		InaccessibleLinks: inAccCount,
	}

	log.Println("link response ", links)
	return links, nil

}

//the ParseTitle function parses the title of the html doc
// breaks with the first occurence of the title tag
func (parser *parserSVC) ParseTitle() (string, error) {
	var title string
	parser.doc.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title = s.Text()
		log.Println("title ->", title)
		return false
	})
	return title, nil
}

//concurrentLinkAccess to concurrenlty access the links
// assumption : the raw url from the html document is used as-is
// if the page is not accessible, add  the inaccessible links counter
func concurrentLinkAccess(link string, wg *sync.WaitGroup, out chan<- int) {

	defer wg.Done()
	var count int
	// converting the raw link to a rquest object for the extractor service to perform Scrape on it
	r := domain.Request{
		URL: link,
	}
	svc := NewExtractorService(r)
	// call the Scrape from the extractor service
	resp, err := svc.Scrape()
	if err != nil {
		count += 1
	} else {
		log.Println("response code for ", link, " ", resp.StatusCode)
		if resp.StatusCode != 200 {
			count += 1
		}
	}
	log.Println("inside concurrent access ", link)
	out <- count

}

// checkChan will monitor the dirrent concurrently running funcions and the channel usage
func checkChan(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
}
