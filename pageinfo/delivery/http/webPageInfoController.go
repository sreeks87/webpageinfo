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

func Heartbeat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, bytes.NewBuffer([]byte(`{"Status":"Ok"}`)))
}

func Webpageinfo(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		HandlePOSTError(http.StatusBadRequest, "request body cant be empty", w)
		return
	}
	req, e := ioutil.ReadAll(r.Body)
	var resp domain.Pageinfo
	if len(req) == 0 || req == nil {

		HandlePOSTError(http.StatusBadRequest, "request is nil or empty", w)
		return
	}
	if e != nil {

		HandlePOSTError(http.StatusBadRequest, e.Error(), w)
		return
	}
	var requestObj domain.Request
	json.Unmarshal(req, &requestObj)
	exSvc := svc.NewExtractorService(requestObj)
	resp, err := exSvc.Extract()
	if err != nil {

		HandlePOSTError(http.StatusBadRequest, err.Error(), w)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}

func HandlePOSTError(status int, e string, w http.ResponseWriter) {
	resp := domain.Pageinfo{
		Error: e,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&resp)
}
