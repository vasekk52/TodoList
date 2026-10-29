package service

import (
	"errors"
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

func TestCreateTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{createTodo: func(title string) (model.Todo, error) {
			if title != "test" {
				t.Fatalf("want: test, got: %v", title)
			}
			return model.Todo{ID: 1, Title: title}, nil
		}}
		service := NewService(svc)

		got, err := service.CreateTodo("test")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.ID != 1 || got.Title != "test" {
			t.Errorf("got: %+v, want ID 1 and title test", got)
		}
	})

	t.Run("empty title", func(t *testing.T) {
		svc := &fakeService{createTodo: func(title string) (model.Todo, error) {
			t.Fatal("service cant called on empty title")
			return model.Todo{}, nil
		}}
		service := NewService(svc)

		_, err := service.CreateTodo("")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("service error", func(t *testing.T) {
		svc := &fakeService{createTodo: func(title string) (model.Todo, error) {
			return model.Todo{}, errors.New("db down")
		}}
		service := NewService(svc)

		_, err := service.CreateTodo("test")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestGetTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{getTodo: func(id int) (model.Todo, error) {
			if id != 2 {
				t.Fatalf("want 2, got: %v", id)
			}
			return model.Todo{ID: 2, Title: "test"}, nil
		}}
		service := NewService(svc)

		got, err := service.GetTodo(2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.ID != 2 || got.Title != "test" {
			t.Errorf("got: %+v, want ID 2 and title test", got)
		}
	})

	t.Run("not found", func(t *testing.T) {
		svc := &fakeService{getTodo: func(id int) (model.Todo, error) {
			return model.Todo{}, errors.New("not found")
		}}
		service := NewService(svc)

		_, err := service.GetTodo(999)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestCompleteTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{completeTodo: func(id int) (model.Todo, error) {
			if id != 3 {
				t.Fatalf("want 3, got: %v", id)
			}
			return model.Todo{ID: 3, Title: "test", Completed: true}, nil
		}}
		service := NewService(svc)

		got, err := service.CompleteTodo(3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Completed {
			t.Errorf("got completed %v, want true", got.Completed)
		}
	})

	t.Run("service error", func(t *testing.T) {
		svc := &fakeService{completeTodo: func(id int) (model.Todo, error) {
			return model.Todo{}, errors.New("update failed")
		}}
		service := NewService(svc)

		_, err := service.CompleteTodo(3)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestDeleteTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		deletedID := 0
		svc := &fakeService{deleteTodo: func(id int) error {
			deletedID = id
			return nil
		}}
		service := NewService(svc)

		if err := service.DeleteTodo(4); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if deletedID != 4 {
			t.Errorf("got id %v, want 4", deletedID)
		}
	})

	t.Run("not found", func(t *testing.T) {
		svc := &fakeService{deleteTodo: func(id int) error {
			return errors.New("not found")
		}}
		service := NewService(svc)

		if err := service.DeleteTodo(999); err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestGetAllTodo(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := &fakeService{getAllTodo: func() ([]model.Todo, error) {
			return []model.Todo{
				{ID: 1, Title: "test 1"},
				{ID: 2, Title: "test 2"},
			}, nil
		}}
		service := NewService(svc)

		got, err := service.GetAllTodo()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got) != 2 {
			t.Fatalf("got len %v, want 2", len(got))
		}
		if got[0].Title != "test 1" || got[1].Title != "test 2" {
			t.Errorf("got: %+v, want valid items", got)
		}
	})

	t.Run("service error", func(t *testing.T) {
		svc := &fakeService{getAllTodo: func() ([]model.Todo, error) {
			return nil, errors.New("db error")
		}}
		service := NewService(svc)

		_, err := service.GetAllTodo()
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}