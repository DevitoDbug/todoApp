package routes

import (
	"github.com/DevitoDbug/todoApp/pkg/controllers"
	"github.com/gorilla/mux"
)

var TaskRoutes = func(r mux.Router) {
	r.HandleFunc("/task/", controllers.GetAllTasks).Methods("GET")
	r.HandleFunc("/task/{id}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/task/", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/task/{id}", controllers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/task/{id}", controllers.UpdateTask).Methods("PUT")
}
