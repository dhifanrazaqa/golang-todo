package routes

import (
	"github.com/gorilla/mux"
	"github.com/dhifanrazaqa/golang-todo/controllers"
)

func Init() *mux.Router {

	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show).Methods("GET", "OPTIONS")
	route.HandleFunc("/add", controllers.Add).Methods("POST", "OPTIONS")
	route.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE", "OPTIONS")
	route.HandleFunc("/complete/{id}", controllers.Complete).Methods("POST", "OPTIONS")

	return route
}