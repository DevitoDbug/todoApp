package utils

import (
	"github.com/gorilla/mux"
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
