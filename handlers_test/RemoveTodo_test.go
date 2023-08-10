package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"

	"github.com/hcastro1515/todosapi/db"
	handlers "github.com/hcastro1515/todosapi/handler"
	"github.com/hcastro1515/todosapi/models"
)

func TestRemoveTodo(t *testing.T) {
	DB := db.Init()
	h := handlers.New(DB)

	todo := models.Todo{
		Title:       "Test Todo",
		Description: "This is a test todo",
	}

	h.DB.Create(&todo)

	req, err := http.NewRequest("DELETE", "/todos/"+strconv.Itoa(int(todo.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(todo.ID))})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.RemoveTodo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var fetchedTodo models.Todo
	if result := h.DB.First(&fetchedTodo, todo.ID); result.Error == nil {
		t.Errorf("handler did not delete todo")
	}
}
