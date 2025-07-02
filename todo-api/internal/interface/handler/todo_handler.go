package handler

import (
	"net/http"
	"strconv"
	"todo-app/internal/usecase"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	usecase *usecase.TodoUsecase
}

func NewTodoHandler(u *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{usecase: u}
}

type createRequest struct {
	Title string `json:"title" validate:"required"`
}

type updateRequest struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed"`
}

func (h *TodoHandler) GetTodos(c echo.Context) error {
	todos, err := h.usecase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.usecase.GetByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Todo not found"})
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	req := new(createRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	todo, err := h.usecase.Create(c.Request().Context(), req.Title)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(updateRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	todo, err := h.usecase.Update(c.Request().Context(), uint(id), req.Title, req.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
