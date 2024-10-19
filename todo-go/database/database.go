package database

import (
	"encoding/json"
	"os"

	"github.com/milanmlft/todo/todo-go/todo"
)

func InitialiseDB(path string) error {
	encoding, err := json.Marshal([]string{})
	if err != nil {
		return err
	}
	return os.WriteFile(path, encoding, 0644)
}

type DatabaseHandler struct {
	path string
}

func (db *DatabaseHandler) ReadTodos() (todo.Todos, error) {
	contents, err := os.ReadFile(db.path)
	if err != nil {
		return todo.Todos{}, err
	}

	var todos todo.Todos
	err = json.Unmarshal(contents, &todos)
	if err != nil {
		return todo.Todos{}, err
	}
	return todos, nil
}

func (db *DatabaseHandler) WriteTodos(todoList todo.Todos) error {
	encoding, err := json.Marshal(todoList)
	if err != nil {
		return err
	}
	return os.WriteFile(db.path, encoding, 0644)
}
