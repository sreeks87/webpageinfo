package service

import (
	"log"
	u "net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
)

const PASSWORD = "password"

func parsePage(doc *goquery.Document, url string) (domain.Pageinfo, error) {

	head, e := parseHead(doc)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	links, e := parseLinks(doc, url)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	title, e := parseTitle(doc)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	version, e := parseHtmlVersion(doc)
	if e != nil {
		return domain.Pageinfo{}, e
	}
	login, e := parseLoginForm(doc)
	if e != nil {
		return domain.Pageinfo{}, e
	}

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

// <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
// <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN""http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
// common part --> -//W3C//DTD [----]//EN
func parseHtmlVersion(doc *goquery.Document) (string, error) {
	commonStart := "-//W3C//DTD"
	commonEnd := "//EN"
	version := "HTML 5.0"
	if len(doc.Selection.Nodes) > 0 {
		if len(doc.Selection.Nodes[0].FirstChild.Attr) > 0 {
			vString := doc.Selection.Nodes[0].FirstChild.Attr[0].Val
			version = strings.Replace(vString, commonStart, "", 1)
			version = strings.Replace(version, commonEnd, "", 1)
		}
	}
	log.Println(doc.Children())
	return version, nil
}
func parseLoginForm(doc *goquery.Document) (bool, error) {
	found := false
	doc.Find("input").EachWithBreak(func(i int, s *goquery.Selection) bool {
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

func parseHead(doc *goquery.Document) (domain.Head, error) {
	head := map[string]int{"h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
	for k, _ := range head {
		log.Println("-------------------------HEADING ", k, "---------------------")
		doc.Find(k).Each(func(i int, s *goquery.Selection) {
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
func parseLinks(doc *goquery.Document, url string) (domain.Links, error) {
	log.Println("-------------------------Links---------------------")
	parsedURL, _ := u.Parse(url)
	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	var exLink int
	var inLink int
	urlSet := make(map[string]bool)
	var inAccessLink = make(chan int)
	// pattern := fmt.Sprintf("[%s]?", url)
	// r, _ := regexp.Compile(pattern)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Attr("href")
		// assumption made : an internal link contains either the url as the starting part or /
		// https://www.google.com/contact and /blog will be considered as internal link for url ->www.google.com
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
			}
		}
	})

	log.Println("internal links count ", inLink)
	log.Println("external links count ", exLink)
	log.Println("links to access ", urlSet)

	var wg sync.WaitGroup
	for link, _ := range urlSet {
		wg.Add(1)
		go concurrentLinkAccess(link, &wg, inAccessLink)
	}

	go checkChan(&wg, inAccessLink)
	links := domain.Links{
		InternalLinks:     inLink,
		ExternalLinks:     exLink,
		InaccessibleLinks: <-inAccessLink,
	}

	return links, nil

}

func parseTitle(doc *goquery.Document) (string, error) {
	var title string
	doc.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title = s.Text()
		log.Println("title ->", title)
		return false
	})
	return title, nil
}

func concurrentLinkAccess(link string, wg *sync.WaitGroup, out chan<- int) {

	defer wg.Done()
	var count int
	resp, err := scrape(link)
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

func checkChan(wg *sync.WaitGroup, ch chan int) {
	wg.Wait()
	close(ch)
}
