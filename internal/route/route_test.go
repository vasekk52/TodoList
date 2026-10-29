package route

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp/internal/handler"
	"todoapp/internal/model"
)

type routes struct {
	name		string
	method		string
	target 		string
}

type fakeSvc struct{}

func (d *fakeSvc) CreateTodo(title string) (model.Todo, error) {
	return model.Todo{}, nil
}
func (d *fakeSvc) GetTodo(id int) (model.Todo, error) {
	return model.Todo{}, nil
}
func (d *fakeSvc) CompleteTodo(id int) (model.Todo, error) {
	return model.Todo{}, nil
}
func (d *fakeSvc) DeleteTodo(id int) error {
	return nil
}
func (d *fakeSvc) GetAllTodo() ([]model.Todo, error) {
	return nil, nil
}
func TestRoute(t *testing.T) {
	svc := &fakeSvc{}
	h := handler.NewHandler(svc)
	r := New(h)

	routes := []routes{
		{"create todo", http.MethodPost, "/todos"},
		{"get todo", http.MethodGet, "/todos"},
		{"get all todos", http.MethodGet, "/todos/1"},
		{"complete todo", http.MethodPatch, "/todos/1"},
		{"delete todo", http.MethodDelete, "/todos/1"},
	}

	for _, route := range routes {
		t.Run(route.name, func(t *testing.T) {
			req := httptest.NewRequest(route.method, route.target, nil)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code == http.StatusNotFound {
				t.Errorf("route %v %v not found, got 404", route.method, route.target)
			}
		})
	}

	t.Run("unreg route 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/123test", nil)
		rec := httptest.NewRecorder()

		r.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("want 404, got %v", rec.Code)
		}
	})
}