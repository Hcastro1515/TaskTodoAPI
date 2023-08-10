package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hcastro1515/todosapi/db"
	handlers "github.com/hcastro1515/todosapi/handler"
	"github.com/hcastro1515/todosapi/models"
)

func TestCreateTodo(t *testing.T) {
	DB := db.Init()
	h := handlers.New(DB)

	todo := models.Todo{
		Title:       "Test Todo",
		Description: "This is a test todo",
		Status:      false,
	}

	body, err := json.Marshal(todo)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var createdTodo models.Todo
	err = json.NewDecoder(rr.Body).Decode(&createdTodo)
	if err != nil {
		t.Fatal(err)
	}

	if createdTodo.Title != todo.Title {
		t.Errorf("handler returned unexpected title: got %v want %v",
			createdTodo.Title, todo.Title)
	}

	if createdTodo.Description != todo.Description {
		t.Errorf("handler returned unexpected description: got %v want %v",
			createdTodo.Description, todo.Description)
	}
}
