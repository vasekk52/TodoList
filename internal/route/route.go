package route

import (
	"net/http"
	"todoapp/internal/handler"
)

func New(todoHandlers *handler.Handler) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("POST /todos", todoHandlers.CreateTodoHandler)
	router.HandleFunc("GET /todos", todoHandlers.GetAllTodoHandler)
	router.HandleFunc("GET /todos/{id}", todoHandlers.GetTodoHandler)
	router.HandleFunc("DELETE /todos/{id}", todoHandlers.DeleteTodoHandler)
	router.HandleFunc("PATCH /todos/{id}", todoHandlers.CompleteTodoHandler)

	return router
}