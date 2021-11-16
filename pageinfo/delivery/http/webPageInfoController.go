package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sreeks87/webpageinfo/pageinfo/domain"
	svc "github.com/sreeks87/webpageinfo/pageinfo/service"
)

// Heartbeat is a heartbeat method used for testing the api
func Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, bytes.NewBuffer([]byte(`{"Status":"Ok"}`)))
}

// Webpageinfo is the the actual webpage information extractor
// the /webpageinfo invokes this function
// returns error when the request is nil/empty
// or if the extractor service returns any error
func Webpageinfo(w http.ResponseWriter, r *http.Request) {
	// if the body itself is nil, then return error
	if r.Body == nil {
		HandlePOSTError(http.StatusBadRequest, "request body cant be empty", w)
		return
	}
	req, e := ioutil.ReadAll(r.Body)
	var resp *domain.Pageinfo
	// checking  if the request body is nil or empty, then return error
	if len(req) == 0 || req == nil {

		HandlePOSTError(http.StatusBadRequest, "request is nil or empty", w)
		return
	}
	// if the readall() fails for any reason the propagate the error to the caller
	if e != nil {

		HandlePOSTError(http.StatusBadRequest, e.Error(), w)
		return
	}
	var requestObj domain.Request
	// parsing the request body to model struct
	json.Unmarshal(req, &requestObj)
	// the extract method is implemented by NewExtractorService object
	// create the object below to call the extractor
	exSvc := svc.NewExtractorService(requestObj)
	resp, err := exSvc.Extract()
	if err != nil {
		HandlePOSTError(http.StatusBadRequest, err.Error(), w)
		return
	}

	// write the final response back to the user
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

// HandlePOSTError function to handle post errors
// writes the error into the Pageinfo struct
func HandlePOSTError(status int, e string, w http.ResponseWriter) {
	resp := domain.Pageinfo{
		Error: e,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&resp)
}
