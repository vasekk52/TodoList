package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todoapp/internal/model"
	"todoapp/internal/service"
)
type Handler struct {
	service *service.TodoService
}

func NewHandler(service *service.TodoService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoDTO := model.TodoDTO{}
	
	json.NewDecoder(r.Body).Decode(&todoDTO)
	createdTodo, err := h.service.CreateTodo(todoDTO.Title)
		if err != nil {
		writeError(w, http.StatusBadRequest, "" + err.Error() )
		return
	}
	writeJSON(w, http.StatusCreated, createdTodo)
}

func (h *Handler) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0{
		writeJSON(w, http.StatusBadRequest, "invalid id")
		return
	}
	
	todo, err := h.service.GetTodo(int(id))
	if err != nil {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	writeJSON(w, 200, todo)

}

func (h *Handler) CompleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0{
		writeJSON(w, http.StatusBadRequest, "invalid id")
		return
	}

	todo, err := h.service.CompleteTodo(int(id))
	if err != nil {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	writeJSON(w, 200, todo)
	
}

func (h *Handler) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0{
		writeJSON(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.service.DeleteTodo(int(id))
	if err != nil {
		writeError(w, http.StatusNotFound, "todo not found")
		return
	}

	writeJSON(w, 200, "")
}

func (h *Handler) GetAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAllTodo()
	if err != nil {
		writeError(w, http.StatusNotFound, "todos not found")
		return
	}
	
	writeJSON(w, 200, todos)
}