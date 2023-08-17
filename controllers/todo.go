package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"github.com/dhifanrazaqa/golang-todo/config"
	"github.com/dhifanrazaqa/golang-todo/models"
	"github.com/gorilla/mux"
)

var (
	id        int
	item      string
	completed int
	database  = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM todos`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &item, &completed)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	todoJson, err := json.Marshal(map[string][]models.Todo{"todo": todos})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(todoJson)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var addTodo struct {
		Item string
	}

	err := json.NewDecoder(r.Body).Decode(&addTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = database.Exec(`INSERT INTO todos (item) VALUE (?)`, addTodo.Item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Todo Added Successfully!"}`))
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]

	_, err := database.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Todo Complete!"}`))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]

	_, err := database.Exec(`DELETE FROM todos WHERE id = ?`, todoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Todo Deleted Successfully!"}`))
}