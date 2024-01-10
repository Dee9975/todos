package store

import (
	"awesomeProject/data"
	"awesomeProject/util"
	"database/sql"
	"errors"
)

type TodoStore struct {
	db *sql.DB
}

func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{
		db: db,
	}
}

func (s *TodoStore) GetTodos() ([]data.TodoResponse, error) {
	rows, err := s.db.Query("select * from todos")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]data.TodoResponse, 0)

	for rows.Next() {
		todo := data.Todo{}
		err = rows.Scan(&todo.Id, &todo.Name, &todo.Complete)

		if err != nil {
			return nil, err
		}
		response := data.TodoResponse{
			Id:       todo.Id,
			Name:     todo.Name,
			Complete: util.IntToBool(todo.Complete),
		}
		todos = append(todos, response)
	}
	err = rows.Err()

	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoStore) FindTodo(id int) (*data.Todo, error) {
	var todo data.Todo
	rows, err := s.db.Query("select * from todos where id=$1 limit 1", id)
	if err != nil {
		return nil, err
	}
	i := 0
	for rows.Next() {
		i++
		if err := rows.Scan(&todo.Id, &todo.Name, &todo.Complete); err != nil {
			return nil, err
		}
	}

	if i == 0 {
		return nil, errors.New("No rows found")
	}

	if rows.Err() != nil {
		return nil, err
	}
	return &todo, nil
}
func (s *TodoStore) CreateTodo(todo data.CreateTodoRequest) error {
	statement := `insert into todos (name, complete) values ($1, $2)`

	_, err := s.db.Exec(statement, todo.Name, todo.Complete)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoStore) UpdateTodo(id int, complete bool) error {
	_, err := s.db.Exec("update todos set complete=$1 where id=$2", util.BoolToInt(complete), id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoStore) DeleteTodo(id int) error {
	_, err := s.db.Exec("delete from todos where id=$1", id)
	return err
}
