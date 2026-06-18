package repository

import (
	"REST-api/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, todo_description string,  completed bool) ( *models.Todo , error ) {

	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query string = `
		INSERT INTO todos (title, todo_description, completed)
		VALUES ($1, $2, $3)
		RETURNING id, title, todo_description, completed, created_at, updated_at
	`

	var todo models.Todo 

	var err error = pool.QueryRow(ctx, query, title, todo_description, completed).Scan(
		&todo.ID,
		&todo.Title,
		&todo.TodoDescription,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}

	return &todo, nil

}