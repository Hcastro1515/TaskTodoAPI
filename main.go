package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hcastro1515/todosapi/db"
	handlers "github.com/hcastro1515/todosapi/handler"
	"github.com/rs/cors"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/todos", h.CreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todos", h.GetTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", h.GetTodoById).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", h.UpdateTodo).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", h.RemoveTodo).Methods(http.MethodDelete)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	handlecors := c.Handler(router)

	fmt.Println("Listening on port 4000")
	http.ListenAndServe(":4000", handlecors)
}
