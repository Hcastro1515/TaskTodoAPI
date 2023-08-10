package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hcastro1515/todosapi/models"
)

func (h Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	w.Header().Add("Content-Type", "application/json")
	h.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}
