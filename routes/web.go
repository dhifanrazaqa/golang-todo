package routes

import (
	"github.com/gorilla/mux"
	"github.com/dhifanrazaqa/golang-todo/controllers"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show).Methods("GET")
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	route.HandleFunc("/complete/{id}", controllers.Complete).Methods("POST")

	return route
}