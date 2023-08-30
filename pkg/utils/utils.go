package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func IDFromRouteVariable(r *http.Request) (ID int64) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		log.Printf("Retrieving id from route error:\n%v\n", err)
		return
	}
	return ID
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			log.Printf("Pursing body error:\n%v\n", err)
			return
		}
	}
}
