# Webpageinfo
The app takes a webpage url and displays information about it.
A working version can be found here - link_to_page

## What the app does

The app expects a valid url to be provided in the search bar.

It then extracts the details based on assumption as mentioed below.

1. HTML Version : The html version of the document is extracted based on
https://www.w3.org/QA/2002/04/valid-dtd-list.html, useful for the older versions of HTML.

    ###  Assumptions : 

    The defaultversion is HTML version 5.

    For any lowerversion, the root html node will have the doctype in the below formats

    `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`

    `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN""http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`

    The common part of these version strings is as below

    `-//W3C//DTD [----]//EN`

    HTML version will be eextracted based o this information

2. Page Title : Extracts the page title, text inside the title tag.

3. Heading count by level : Extracts the count of h1 - h6 headings in the html

4. Amount of internal,external, inaacessible links

    ### Assumption : 
    
    1. A link is considered internal if it starts in the href with "/" or the same basepath of url
    2. A link is considered external if the base path do not match and has http as the prefix, this was done because some html pages were seen with comment ids as links and these were getting treated as external liks
    3. To check if the link is accessible : make a get call to the link concurrenlty.
    The input to this concurrent function will be the links map urlSet to make sure the links are not duplicated in the input to the function.

5. Presence of  a login form.

    ### Assumption : 
    1. The login form will contain an input of type password 
    If yes, then login form present == true, else false
    
    The app has been built in go, version `1.16`

# To run this app locally

Prerequisite
1. Golang installation

Run locally

1. git clone the repository

2. Download the dependencies `go mod dowmload`

3. Run `go run .`

4. Navigate to `http://localhost:8080/ui/ `

5. Port can be modified in the config file

### Enhancements
1. 

