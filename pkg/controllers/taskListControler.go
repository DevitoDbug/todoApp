package controllers

import (
	"encoding/json"
	"github.com/DevitoDbug/todoApp/pkg/models"
	"github.com/DevitoDbug/todoApp/pkg/utils"
	"log"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()

	res, err := json.Marshal(tasks)
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
	ID := utils.IDFromRouteVariable(r)

	searchedTask, dbInfo := models.GetTaskBYID(ID)
	if dbInfo.Error != nil {
		log.Printf("Database error:\n%v\n", dbInfo.Error)
		return
	}

	res, err := json.Marshal(searchedTask)
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

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask = &models.Task{}
	utils.ParseBody(r, newTask)
	createdTask := newTask.CreateTask()

	res, err := json.Marshal(createdTask)
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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	ID := utils.IDFromRouteVariable(r)
	deletedTask := models.DeleteTask(ID)

	res, err := json.Marshal(deletedTask)
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	newTask := &models.Task{}
	ID := utils.IDFromRouteVariable(r)
	utils.ParseBody(r, newTask)

	updateTask, dbInfo := models.GetTaskBYID(ID)
	if dbInfo.Error != nil {
		log.Printf("Getting data from DB error:\n%v\n", dbInfo.Error)
		return
	}
	if newTask.Name != "" {
		updateTask.Name = newTask.Name
	}
	if newTask.Status != "" {
		updateTask.Status = newTask.Status
	}
	if newTask.Priority != "" {
		updateTask.Priority = newTask.Priority
	}
	if newTask.Notes != "" {
		updateTask.Notes = newTask.Notes
	}

	//updating the db
	dbInfo.Save(&updateTask)

	res, err := json.Marshal(updateTask)
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
