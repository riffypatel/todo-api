package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riffypatel/todo-api/models"
)

var todos []models.Todo

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var updated models.Todo
	json.NewDecoder(r.Body).Decode(&updated)

	for i, t := range todos {
		if t.ID == id {
			todos[i].Title = updated.Title
			todos[i].Done = updated.Done
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	Params := mux.Vars(r)
	id, _ := strconv.Atoi(Params["id"])

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}
