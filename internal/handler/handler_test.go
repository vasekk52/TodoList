package handler

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoapp/internal/model"
)

type fakeService struct {
	createTodo   func(title string) (model.Todo, error)
	getTodo      func(id int) (model.Todo, error)
	completeTodo func(id int) (model.Todo, error)
	deleteTodo   func(id int) error
	getAllTodo   func() ([]model.Todo, error)
}

func (m *fakeService) CreateTodo(title string) (model.Todo, error) {
	if m.createTodo != nil {
		return m.createTodo(title)
	}
	return model.Todo{}, nil
}

func (m *fakeService) GetTodo(id int) (model.Todo, error) {
	if m.getTodo != nil {
		return m.getTodo(id)
	}
	return model.Todo{}, nil
}

func (m *fakeService) CompleteTodo(id int) (model.Todo, error) {
	if m.completeTodo != nil {
		return m.completeTodo(id)
	}
	return model.Todo{}, nil
}

func (m *fakeService) DeleteTodo(id int) error {
	if m.deleteTodo != nil {
		return m.deleteTodo(id)
	}
	return nil
}

func (m *fakeService) GetAllTodo() ([]model.Todo, error) {
	if m.getAllTodo != nil {
		return m.getAllTodo()
	}
	return nil, nil
}

func TestCreateTodoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{createTodo: func(title string) (model.Todo, error) {
			if title != "test" {
				t.Fatalf("got: %v, want: `test`", title)
			}
			return model.Todo{ID: 1, Title: title, Completed: false}, nil
		}}
		h := NewHandler(svc)

		body := strings.NewReader(`{"title": "test"}`)
		req := httptest.NewRequest(http.MethodPost, "/todos", body)
		rec := httptest.NewRecorder()

		h.CreateTodoHandler(rec, req)

		if rec.Code != http.StatusCreated {
			t.Errorf("got: %v, want: %v", rec.Code, http.StatusCreated)
		}
	})

	t.Run("service error", func (t *testing.T)  {
		svc := &fakeService{createTodo: func(title string) (model.Todo, error) {
			return model.Todo{}, errors.New("title cant be empty")
		}}
		h := NewHandler(svc)

		body := strings.NewReader(`{"title: ""}`)
		req := httptest.NewRequest(http.MethodPost, "/todos", body)
		rec := httptest.NewRecorder()

		h.CreateTodoHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("got: %v, want: %v", rec.Code, http.StatusBadRequest)
		}
	})
}

func TestGetTodoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{getTodo: func(id int) (model.Todo, error) {
			if id != 10 {
				t.Fatalf("want: 10, got: %v", id)
			}
			return model.Todo{ID: 10, Title: "test"}, nil
		}}
		h := NewHandler(svc)

		req := httptest.NewRequest(http.MethodGet, "/todos/10", nil)
		req.SetPathValue("id", "10")
		rec := httptest.NewRecorder()

		h.GetTodoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusOK)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		ids := []string{"abc", "0", "-5", ""}

		for _, id := range ids {
		h := NewHandler(&fakeService{})
		
		req := httptest.NewRequest(http.MethodGet, "/todos/"+id, nil)
		req.SetPathValue("id", id)
		rec := httptest.NewRecorder()

		h.GetTodoHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusBadRequest)
		}
	}
	})

	t.Run("not found", func(t *testing.T) {
		svc := &fakeService{getTodo: func(id int) (model.Todo, error) {
			return model.Todo{}, errors.New("todo not found")
		}}
		h := NewHandler(svc)

		req := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		req.SetPathValue("id", "1")
		rec := httptest.NewRecorder()

		h.GetTodoHandler(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusNotFound)
		}
	})	
}
func TestCompleteTodoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{completeTodo: func(id int) (model.Todo, error) {
			if id != 1 {
				t.Fatalf("want 1, got %v", id)
			}
			return model.Todo{ID: 1, Title: "test", Completed: true}, nil
		}}
		h := NewHandler(svc)
		req := httptest.NewRequest(http.MethodPatch, "/todos/1", nil)
		req.SetPathValue("id", "1")
		rec := httptest.NewRecorder()

		h.CompleteTodoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusOK)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		ids := []string{"abc", "0", "-5", ""}

		for _, id := range ids {
		h := NewHandler(&fakeService{})
		
		req := httptest.NewRequest(http.MethodPatch, "/todos/"+id, nil)
		req.SetPathValue("id", id)
		rec := httptest.NewRecorder()

		h.CompleteTodoHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusBadRequest)
		}
	}
	})

	t.Run("not found", func(t *testing.T) {
		svc := &fakeService{completeTodo: func(id int) (model.Todo, error) {
				return model.Todo{}, errors.New("todo not found")
			},
		}

		h := NewHandler(svc)
		req := httptest.NewRequest(http.MethodPatch, "/todos/5", nil)
		req.SetPathValue("id", "5")
		rec := httptest.NewRecorder()

		h.CompleteTodoHandler(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusNotFound)
		}
	})
}

func TestDeleteTodoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deletedID := 0
		svc := &fakeService{deleteTodo: func(id int) error {
			deletedID = id
			return nil
		}}
		h := NewHandler(svc)
		req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
		req.SetPathValue("id", "7")
		rec := httptest.NewRecorder()

		h.DeleteTodoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusOK)
		}
		if deletedID != 7 {
			t.Errorf("got id %v, want id 7", deletedID)
		}
	})
	
	t.Run("invalid id", func(t *testing.T) {
		ids := []string{"abc", "0", "-5", ""}

		for _, id := range ids {
		h := NewHandler(&fakeService{})
		
		req := httptest.NewRequest(http.MethodDelete, "/todos/"+id, nil)
		req.SetPathValue("id", id)
		rec := httptest.NewRecorder()

		h.DeleteTodoHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusBadRequest)
		}
	}
	})

	t.Run("not found", func(t *testing.T) {
		svc := &fakeService{
			deleteTodo: func(id int) error {
				return errors.New("todo not found")
			},
		}

		h := NewHandler(svc)
		req := httptest.NewRequest(http.MethodDelete, "/todos/8", nil)
		req.SetPathValue("id", "8")
		rec := httptest.NewRecorder()

		h.DeleteTodoHandler(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("got status %v, want %v", rec.Code, http.StatusNotFound)
		}
	})
}

func TestGetAllTodoHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{getAllTodo: func() ([]model.Todo, error) {
			return []model.Todo{
				{ID: 1, Title: "test"},
				{ID: 2, Title: "test"},
			}, nil
		}}


		h := NewHandler(svc)
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec := httptest.NewRecorder()

		h.GetAllTodoHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("got status %v, got %v", rec.Code, http.StatusOK)
		}
		if !strings.Contains(rec.Body.String(), "test") {
			t.Errorf("want 'test', got %v", rec.Body.String())
		}
	})

	t.Run("service error", func(t *testing.T) {
		svc := &fakeService{
			getAllTodo: func() ([]model.Todo, error) {
				return nil, errors.New("db error")
			},
		}

		h := NewHandler(svc)
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec := httptest.NewRecorder()

		h.GetAllTodoHandler(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, rec.Code)
		}
	})	
}

func TestWriteJson(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		data := map[string]string{"message":"ok"}

		writeJSON(rec, http.StatusOK, data)

		if rec.Code != http.StatusOK {
			t.Errorf("got %v, want %v", rec.Code, http.StatusOK)
		}

		expectedCT := "application/json; charset=utf-8"
		if ct := rec.Header().Get("Content-Type"); ct != expectedCT {
			t.Errorf("got ct %v, want %v", ct, expectedCT)
		}

		var res map[string]string
		if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
			t.Fatalf("error decode body: %v", err)
		}
		if res["message"] != "ok" {
			t.Errorf("got %v, want %v", res["message"], "ok")
		}

	})

	t.Run("error", func(t *testing.T) {
		rec := httptest.NewRecorder()

		invalidJsonData := math.NaN()

		writeJSON(rec, http.StatusOK, invalidJsonData)
		if rec.Code != http.StatusOK {
			t.Errorf("got %v, want %v", rec.Code, http.StatusOK)
		}
	})
}

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()

	writeError(rec, http.StatusBadRequest, "invalid payload")

	if rec.Code != http.StatusBadRequest {
		t.Errorf("got %v, want %v", rec.Code, http.StatusBadRequest)
	}

	expectedCT := "application/json; charset=utf-8"
	if ct := rec.Header().Get("Content-Type"); ct != expectedCT {
		t.Errorf("got Content-Type %q, want %q", ct, expectedCT)
	}

	var res ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("failed decode error body: %v", err)
	}

	if res.Error != "invalid payload" {
		t.Errorf("got %v, want %v", res.Error, "invalid payload")
	}
}