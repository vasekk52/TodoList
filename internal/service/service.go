package service

import (
	"todoapp/internal/database"
	"todoapp/internal/model"
)

type TodoService struct {
	db *database.DB
}

func NewService(db *database.DB) *TodoService {
	return &TodoService{db:db}
}

func (s *TodoService) CreateTodo(title string) (model.Todo, error){
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