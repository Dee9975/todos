package handlers

import "awesomeProject/store"

type Handler struct {
	todoStore *store.TodoStore
}

func NewHandler(t *store.TodoStore) *Handler {
	return &Handler{
		todoStore: t,
	}
}
