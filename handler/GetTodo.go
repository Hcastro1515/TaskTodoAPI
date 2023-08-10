package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/todosapi/models"
)

func (h Handler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var todo models.Todo
	if result := h.DB.First(&todo, id); result.Error != nil {
		log.Fatal(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}
