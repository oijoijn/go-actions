package repository

import (
	"context"
	"todo-app/internal/domain/model"
)

// TodoRepository はTodoデータの永続化に関するインターフェースです。
type TodoRepository interface {
	FindAll(ctx context.Context) ([]model.Todo, error)
	FindByID(ctx context.Context, id uint) (*model.Todo, error)
	Create(ctx context.Context, todo *model.Todo) error
	Update(ctx context.Context, todo *model.Todo) error
	Delete(ctx context.Context, id uint) error
}
