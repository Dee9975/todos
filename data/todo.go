package data

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Complete int    `json:"complete"`
}

type CreateTodoRequest struct {
	Name     string `json:"name"`
	Complete int    `json:"complete"`
}

type TodoResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

type UpdateTodoRequest struct {
	Complete bool `json:"complete"`
}
