package service

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/sreeks87/webpageinfo/pageinfo/domain"
)

func parsePage(doc *goquery.Document, url string) (domain.Pageinfo, error) {

	go parseHead(doc)
	go parseLinks(doc, url)
	go parseTitle(doc)

	return domain.Pageinfo{}, nil
}

func parseHead(doc *goquery.Document) (domain.Head, error) {
	head := map[string]int{"h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
	for k, v := range head {
		fmt.Println("-------------------------HEADING ", v, "---------------------")
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
	fmt.Println("-------------------------Links---------------------")
	var exLink int
	var inLink int
	var urlSlice []string
	var inAccessLink = make(chan int)
	// pattern := fmt.Sprintf("[%s]?", url)
	// r, _ := regexp.Compile(pattern)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		t, _ := s.Attr("href")
		// assumption made : an internal link contains either the url as the starting part or /
		// https://www.google.com/contact and /blog will be considered as internal link for url ->www.google.com
		if strings.HasPrefix(t, url) {
			fmt.Println("internal link ->", t)
			inLink += 1
			urlSlice = append(urlSlice, url)
		} else if strings.HasPrefix(t, "/") {
			fmt.Println("internal link without url->", t)
			inLink += 1
			urlSlice = append(urlSlice, url+t)
		} else {
			fmt.Println("external link ->", t)
			exLink += 1
		}
	})

	fmt.Println("internal links count ", inLink)
	fmt.Println("external links count ", exLink)
	fmt.Println("links to access ", urlSlice)

	var wg sync.WaitGroup
	for _, link := range urlSlice {
		wg.Add(1)
		go concurrentLinkAccess(link, &wg, inAccessLink)
	}

	wg.Wait()
	close(inAccessLink)
	links := domain.Links{
		InternalLinks:     inLink,
		ExternalLinks:     exLink,
		InaccessibleLinks: <-inAccessLink,
	}

	return links, nil

}

func parseTitle(doc *goquery.Document) (string, error) {
	var title string
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
		fmt.Println("title ->", title)
	})
	return title, nil
}

func concurrentLinkAccess(link string, wg *sync.WaitGroup, out chan<- int) {
	defer wg.Done()
	var count int
	resp, err := scrape(link)
	if err != nil {
		count += 1
	}
	if resp.StatusCode != 200 {
		count += 1
	}
	out <- count
}
