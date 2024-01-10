package handlers

import (
	"awesomeProject/data"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) HandleGetTodos(c echo.Context) error {
	todos, err := h.todoStore.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) HandleCreateTodo(c echo.Context) error {
	var todo data.CreateTodoRequest
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error parsing request body"})
	}
	err := h.todoStore.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed creating todo"})
	}

	return c.JSON(http.StatusCreated, todo)
}

func (h *Handler) HandleUpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var complete data.UpdateTodoRequest
	if err := c.Bind(&complete); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Error parsing body"})
	}
	err := h.todoStore.UpdateTodo(id, complete.Complete)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error updating the todo"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}

func (h *Handler) HandleDeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.todoStore.DeleteTodo(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error updating the todo"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}

func (h *Handler) HandleFindTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.todoStore.FindTodo(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Couldn't find the todo"})
	}

	return c.JSON(http.StatusOK, todo)
}
