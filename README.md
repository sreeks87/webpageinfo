# Webpageinfo
The app takes a webpage url and displays information about it.
A working version can be found here - https://webpageinfo.eu-gb.mybluemix.net/ui/ 

## What the app does

The app expects a valid url to be provided in the search bar.

It then extracts the details based on assumption as mentioed below.

1. HTML Version : The html version of the document is extracted based on
https://www.w3.org/QA/2002/04/valid-dtd-list.html, useful for the older versions of HTML.

    ###  Assumptions : 

    The default version is HTML version 5.

    For any lower version, the root html node will have the doctype in the below formats

    `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">`

    `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN""http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">`

    The common part of these version strings is as below

    `-//W3C//DTD [----]//EN`

    HTML version will be extracted based on this information

2. Page Title : Extracts the page title, the text inside the title tag.

3. Heading count by level : Extracts the count of h1 - h6 headings in the html

4. Amount of internal,external, inaccessible links

    ### Assumptions : 
    
    1. A link is considered internal if it starts in the href tag with "/" or the same basepath of url

    2. A link is considered external if the base path do not match and has http as the prefix, this was done because some html pages were seen with comment ids as links and these were getting treated as external liks

    3. To check if the link is accessible : make a get call to the link concurrenlty.
    The input to this concurrent function will be the map of links to make sure the links are not duplicated in the input to the function.

5. Presence of  a login form.

    ### Assumptions : 
    1. The login form will contain an input of type password. 
    If yes, then login form present == true, else false
    

# To run this app locally

The app has been built in go, version `1.16`

Prerequisite
1. Golang installation

Run locally

1. `git clone the repository`

2. Download the dependencies `go mod download`

3. Run `go run .`

4. Navigate to `http://localhost:8080/ui/ `

5. Port can be modified in the config file


### Design

The application contains 3 layers.

1. The delivery layer, which is the entry point into the application. This deals with the http calls; request and response to and from the system.

2. The service layer : here is where the actual logic of the application resides.
This is where the implementation of the interfaces defined in the domain (below) layer is done.

3. Domain layer, this defines the domain model and any interfaces being used.

The application makes external call to the URL provided by the user, this is therefore defined as an Interface called Service. 

The service interface is responsible for three functions

1. Validate the request from the user

2. Scrape the URL to get the html document

3. Extract the html details from the document

The extraction function inside the Service further depends on the Parser. The parser used for making it easy to parse the html is goquery. Therefore it has also been defined as an interface, if tommorow we decide to use a different parser than this, we could do so by simply writing another parser implementation. This also makes testing easier.

- A simple html page with JS serves as the frontend UI. When the user inputs the URL and clicks the search button, a POST call is made from UI to the backend API.

### API Details

        POST http://localhost:8080/webpageinfo

        {
            
        "url":"http://localhost:8080/ui/"
        
        }

- Unit tests and mocks have been written to cover most of the code

### Deployment and hosting:
    
The app has been hosted on IBM CloudFoundry and the link to the application is https://webpageinfo.eu-gb.mybluemix.net/ui/

### Problems with the implementation and Further Enhancements 

1. Ideally, when there is no object being created the POST method is not preferred, but in this case the method on the endpoint that does the extraction is defined as POST, this is an **assumption** made.

2. The http client is coupled with the Service interface, ideally this could be also a different interface that implements the methods we need to use. 

3. Currently, there is no timeout defined in the http call. But it would be ideal to keep a timeout for the external calls, so that we could treat delayed response as inaccessible link.
We could also take a look at retry mechanism if the links are delayed/not responding to conclude if the link is inaccessible or not.


4. To reduce the external calls, a caching mechanism could be implemented, so that we cache the results and serve the response from the cache.

5. A basic logging has been implemented, but we could take a look at better logging mechanisms to log what level of tracing is required for a particular environment. We could also log the messages to a repository to enable easy querying and extraction via tools like splunk, to help us debug the issues.







