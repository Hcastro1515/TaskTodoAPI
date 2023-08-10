package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/hcastro1515/todosapi/db"
	handlers "github.com/hcastro1515/todosapi/handler"
	"github.com/hcastro1515/todosapi/models"
)

func TestGetTodoById(t *testing.T) {
	DB := db.Init()
	h := handlers.New(DB)

	todo := models.Todo{
		Title:       "Test Todo",
		Description: "This is a test todo",
	}

	h.DB.Create(&todo)

	req, err := http.NewRequest("GET", "/todos/"+strconv.Itoa(int(todo.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(todo.ID))})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetTodoById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var fetchedTodo models.Todo
	err = json.NewDecoder(rr.Body).Decode(&fetchedTodo)
	if err != nil {
		t.Fatal(err)
	}

	if fetchedTodo.ID != todo.ID {
		t.Errorf("handler returned unexpected ID: got %v want %v",
			fetchedTodo.ID, todo.ID)
	}

	if fetchedTodo.Title != todo.Title {
		t.Errorf("handler returned unexpected title: got %v want %v",
			fetchedTodo.Title, todo.Title)
	}

	if fetchedTodo.Description != todo.Description {
		t.Errorf("handler returned unexpected description: got %v want %v",
			fetchedTodo.Description, todo.Description)
	}
}
