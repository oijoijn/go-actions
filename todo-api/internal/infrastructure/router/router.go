package router

import (
	"todo-app/internal/interface/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(h *handler.TodoHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	api := e.Group("/api/v1")
	{
		todos := api.Group("/todos")
		todos.GET("", h.GetTodos)
		todos.POST("", h.CreateTodo)
		todos.GET("/:id", h.GetTodo)
		todos.PUT("/:id", h.UpdateTodo)
		todos.DELETE("/:id", h.DeleteTodo)
	}

	return e
}
