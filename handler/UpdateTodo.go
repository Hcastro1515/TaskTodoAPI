package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/todosapi/models"
)

func (h Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var updatedTodo models.Todo
	json.Unmarshal(body, &updatedTodo)

	var todo models.Todo
	if result := h.DB.First(&todo, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description
	todo.Status = updatedTodo.Status

	h.DB.Save(&todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}
