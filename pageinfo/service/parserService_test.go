package service_test

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/sreeks87/webpageinfo/pageinfo/service"
	"github.com/stretchr/testify/assert"
)

// dummy html string for testing
const htmlDoc = `<!DOCTYPE html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
    <script src="index.js"></script>
    <title>Home24 web page information tool</title>
</head>

<body>
    <div class="container">
        <br/>
        <div class="row justify-content-center">
             <h1>Enter a URL below</h1>
        </div>
    </div>

    <!-- Search bar to appear here -->
<!-- make it look good :) -->
<div class="container">
<div class="container">
    <br/>
</div>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
</body>`

// the below function tests the ParsePage method
// it reads the html string above
// and finds the final input to be sent to the client
// expected :the title should be equal to "Home24 web page information tool"
// in the final response object
func TestParsePage(t *testing.T) {
	str := strings.NewReader(htmlDoc)
	d, _ := goquery.NewDocumentFromReader(str)
	p := service.NewParserSvc(d)
	res, _ := p.ParsePage("http://www.google.com")
	assert.Equal(t, res.PageTitle, "Home24 web page information tool")
}

// the below function tests the ParseHtmlVersion method
// it reads the html string above
// and finds the version of the html
// expected :the version should be equal to "HTML 5.0"
func TestParseHtmlVersion(t *testing.T) {
	str := strings.NewReader(htmlDoc)
	d, _ := goquery.NewDocumentFromReader(str)
	p := service.NewParserSvc(d)
	res, _ := p.ParsePage("http://www.google.com")
	assert.Equal(t, res.HTMLVersion, "HTML 5.0")
}

// the below function tests the ParseHead method
// it reads the html string above
// and finds the count of headings h1-h6
// expected :the h1 count in the above html string is 1
func TestParseHead(t *testing.T) {
	str := strings.NewReader(htmlDoc)
	d, _ := goquery.NewDocumentFromReader(str)
	p := service.NewParserSvc(d)
	res, _ := p.ParsePage("http://www.google.com")
	assert.Equal(t, res.HeadingData.H1, 1)
}

// the below function tests the ParseTitle method
// it reads the html string above
// and finds the title of the page
// expected :the title should be "Home24 web page information tool"
func TestParseTitle(t *testing.T) {
	str := strings.NewReader(htmlDoc)
	d, _ := goquery.NewDocumentFromReader(str)
	p := service.NewParserSvc(d)
	res, _ := p.ParsePage("http://www.google.com")
	assert.Equal(t, res.PageTitle, "Home24 web page information tool")
}
