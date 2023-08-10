package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hcastro1515/todosapi/models"
)

func (h Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	if result := h.DB.Find(&todos); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
