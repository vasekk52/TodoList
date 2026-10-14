package database

import (
	"errors"
	"fmt"
	"todoapp/internal/model"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db 	*sqlx.DB
}

func NewPostgresDB(cfg model.Config) (*DB, error) {
	connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME,
    )
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
	  return nil, err
	}
	return &DB{db:db}, nil
}

func (db *DB) CreateTodo(title string) (model.Todo, error){
	todo := model.Todo{}
	
	err := db.db.QueryRow(`
	INSERT INTO todos (title, createdAt)
	VALUES ($1, NOW())
	RETURNING id, title, completed, createdAt`, title).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
	  return model.Todo{}, err
	}
	return todo, nil
}

func (db *DB) GetTodo(id int) (model.Todo, error) {
	todo := model.Todo{}
	err := db.db.QueryRowx(`
	SELECT * FROM todos 
	WHERE id = $1`, id).StructScan(&todo); if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}

func (db *DB) CompleteTodo(id int) (model.Todo, error) {
	todo := model.Todo{}

	err := db.db.QueryRowx(`
	UPDATE todos
	SET completed   = NOT completed,
    	completedAt = CASE
                    	WHEN NOT completed THEN NOW()
                    	ELSE NULL
                  		END
	WHERE id = $1
	RETURNING id, title, completed, createdAt, completedAt`, id).StructScan(&todo); if err != nil {
		return model.Todo{}, err
	}
	
	
	return todo, nil
}

func (db *DB) DeleteTodo(id int) error {
    var deletedID int64

    err := db.db.QueryRowx(`
        DELETE FROM todos
        WHERE id = $1
        RETURNING id`, id).Scan(&deletedID)

    if errors.Is(err, sql.ErrNoRows) {
        return errors.New("not found")
    }
    if err != nil {
        return fmt.Errorf("delete todo: %w", err)
    }

    return nil
}

func (db *DB) GetAllTodo() ([]model.Todo, error) {
	todos := []model.Todo{}
	if err := db.db.Select(&todos,`
	SELECT * FROM todos`); err != nil {
		return todos, err 
	}
	return todos, nil
}