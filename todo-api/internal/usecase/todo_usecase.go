package usecase

import (
	"context"
	"encoding/json"
	"todo-app/internal/domain/model"
	"todo-app/internal/domain/repository"
)

// EventPublisher はイベントを発行するためのインターフェースです。
type EventPublisher interface {
	Publish(ctx context.Context, subject string, data []byte) error
}

// TodoUsecase はTodoに関するビジネスロジックを実装します。
type TodoUsecase struct {
	repo      repository.TodoRepository
	publisher EventPublisher
}

// NewTodoUsecase は新しいTodoUsecaseのインスタンスを生成します。
func NewTodoUsecase(repo repository.TodoRepository, publisher EventPublisher) *TodoUsecase {
	return &TodoUsecase{repo: repo, publisher: publisher}
}

func (u *TodoUsecase) GetAll(ctx context.Context) ([]model.Todo, error) {
	return u.repo.FindAll(ctx)
}

func (u *TodoUsecase) GetByID(ctx context.Context, id uint) (*model.Todo, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *TodoUsecase) Create(ctx context.Context, title string) (*model.Todo, error) {
	todo := &model.Todo{
		Title:     title,
		Completed: false,
	}

	if err := u.repo.Create(ctx, todo); err != nil {
		return nil, err
	}
	
	// イベントを発行
	eventData, _ := json.Marshal(todo)
	_ = u.publisher.Publish(ctx, "TODOS.created", eventData)

	return todo, nil
}

func (u *TodoUsecase) Update(ctx context.Context, id uint, title string, completed bool) (*model.Todo, error) {
	todo, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	todo.Title = title
	todo.Completed = completed

	if err := u.repo.Update(ctx, todo); err != nil {
		return nil, err
	}
	
	// イベントを発行
	eventData, _ := json.Marshal(todo)
	_ = u.publisher.Publish(ctx, "TODOS.updated", eventData)

	return todo, nil
}

func (u *TodoUsecase) Delete(ctx context.Context, id uint) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return err
	}
	
	// イベントを発行
	eventData, _ := json.Marshal(map[string]uint{"id": id})
	_ = u.publisher.Publish(ctx, "TODOS.deleted", eventData)

	return nil
}
