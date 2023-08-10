package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hcastro1515/todosapi/models"
)

func (h Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todos models.Todo

	err := json.NewDecoder(r.Body).Decode(&todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todos.Title == "" || todos.Description == "" {
		http.Error(w, "All fields must be provided", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	if result := h.DB.Create(&todos); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todos)
}
