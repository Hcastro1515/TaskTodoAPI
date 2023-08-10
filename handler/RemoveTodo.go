package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/todosapi/models"
)

func (h Handler) RemoveTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if result := h.DB.Where("id = ?", id).Delete(&models.Todo{}); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Todo Deleted")
}
