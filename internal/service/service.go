package service

import (
	"errors"
	"todoapp/internal/model"
)

type TodoRepository interface {
    CreateTodo(title string) (model.Todo, error)
    GetTodo(id int) (model.Todo, error)
    CompleteTodo(id int) (model.Todo, error)
    DeleteTodo(id int) error
    GetAllTodo() ([]model.Todo, error)
}

type TodoService struct {
	db TodoRepository
}

func NewService(db TodoRepository) *TodoService {
	return &TodoService{db:db}
}

func (s *TodoService) CreateTodo(title string) (model.Todo, error){
	if title == "" {
		return model.Todo{}, errors.New("title cant be empty")
	}
	todo, err := s.db.CreateTodo(title)
	if err != nil {
	  return model.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) GetTodo(id int) (model.Todo, error){
	todo, err := s.db.GetTodo(id)
	if err != nil {
	  return model.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) CompleteTodo(id int) (model.Todo, error){
	todo, err := s.db.CompleteTodo(id)
	if err != nil {
	  return model.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) DeleteTodo(id int) (error){
	err := s.db.DeleteTodo(id)
	if err != nil {
	  return err
	}

	return nil
}

func (s *TodoService) GetAllTodo() ([]model.Todo, error){
	todos, err := s.db.GetAllTodo()
	if err != nil {
	  return todos, err
	}

	return todos, nil
}