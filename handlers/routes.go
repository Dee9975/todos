package handlers

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	v1.GET("/todos", h.HandleGetTodos)
	v1.POST("/todos", h.HandleCreateTodo)
	v1.PUT("/todos/:id", h.HandleUpdateTodo)
	v1.DELETE("/todos/:id", h.HandleDeleteTodo)
	v1.GET("/todos/:id", h.HandleFindTodo)
}
