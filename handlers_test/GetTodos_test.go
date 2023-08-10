package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hcastro1515/todosapi/db"
	handlers "github.com/hcastro1515/todosapi/handler"
	"github.com/hcastro1515/todosapi/models"
)

func TestGetTodos(t *testing.T) {
	DB := db.Init()
	h := handlers.New(DB)

	todo1 := models.Todo{
		Title:       "Test Todo 1",
		Description: "This is a test todo 1",
	}

	todo2 := models.Todo{
		Title:       "Test Todo 2",
		Description: "This is a test todo 2",
	}

	h.DB.Create(&todo1)
	h.DB.Create(&todo2)

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetTodos)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var fetchedTodos []models.Todo
	err = json.NewDecoder(rr.Body).Decode(&fetchedTodos)
	if err != nil {
		t.Fatal(err)
	}
}
