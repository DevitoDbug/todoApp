package controllers

import (
	"encoding/json"
	"github.com/DevitoDbug/todoApp/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()

	res, err := json.Marshal(*tasks)
	if err != nil {
		log.Printf("Could not marshal tasks to json:\n%v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		log.Printf("Writting to header error:\n%v\n", err)
		return
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
}
func CreateTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
