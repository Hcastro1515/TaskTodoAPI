package handlers_test

import (
	"bytes"
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

func TestUpdateTodo(t *testing.T) {
	DB := db.Init()
	h := handlers.New(DB)

	todo := models.Todo{
		Title:       "Test Todo",
		Description: "This is a test todo",
	}

	h.DB.Create(&todo)

	updatedTodo := models.Todo{
		Title:       "Updated Test Todo",
		Description: "This is an updated test todo",
		Status:      true,
	}

	body, err := json.Marshal(updatedTodo)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/todos/"+strconv.Itoa(int(todo.ID)), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(todo.ID))})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.UpdateTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var fetchedTodo models.Todo
	h.DB.First(&fetchedTodo, todo.ID)

	if fetchedTodo.Title != updatedTodo.Title {
		t.Errorf("handler did not update title: got %v want %v",
			fetchedTodo.Title, updatedTodo.Title)
	}

	if fetchedTodo.Description != updatedTodo.Description {
		t.Errorf("handler did not update description: got %v want %v",
			fetchedTodo.Description, updatedTodo.Description)
	}

	if fetchedTodo.Status != updatedTodo.Status {
		t.Errorf("handler did not update status: got %v want %v",
			fetchedTodo.Status, updatedTodo.Status)
	}
}
