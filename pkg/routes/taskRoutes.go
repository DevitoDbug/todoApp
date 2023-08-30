package routes

import (
	"github.com/DevitoDbug/todoApp/pkg/controllers"
	"github.com/gorilla/mux"
)

var TaskRoutes = func(r mux.Router) {
	r.HandleFunc("/task/", controllers.GetAllTasks)
	r.HandleFunc("/task/", controllers.GetTask)
	r.HandleFunc("/task/", controllers.CreateTask)
	r.HandleFunc("/task/", controllers.DeleteTask)
	r.HandleFunc("/task/", controllers.UpdateTask)
}
