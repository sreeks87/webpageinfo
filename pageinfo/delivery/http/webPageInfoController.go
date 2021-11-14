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
	req, e := ioutil.ReadAll(r.Body)
	var resp domain.Pageinfo
	if len(req) == 0 {

		HandlePOSTError(400, e.Error(), w)
		return
	}
	if e != nil {

		HandlePOSTError(400, e.Error(), w)
		return
	}
	var requestObj domain.Request
	json.Unmarshal(req, &requestObj)

	resp, err := svc.Extract(&requestObj)
	if err != nil {

		HandlePOSTError(400, err.Error(), w)
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
