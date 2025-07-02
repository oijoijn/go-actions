package persistence

import (
	"context"
	"todo-app/internal/domain/model"
	"todo-app/internal/domain/repository"

	"gorm.io/gorm"
)

type gormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &gormTodoRepository{db: db}
}

func (r *gormTodoRepository) FindAll(ctx context.Context) ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.WithContext(ctx).Find(&todos).Error
	return todos, err
}

func (r *gormTodoRepository) FindByID(ctx context.Context, id uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.WithContext(ctx).First(&todo, id).Error
	return &todo, err
}

func (r *gormTodoRepository) Create(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *gormTodoRepository) Update(ctx context.Context, todo *model.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *gormTodoRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Todo{}, id).Error
}
