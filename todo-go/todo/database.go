package todo

import (
	"encoding/json"
	"os"
)

func InitialiseDB(path string) error {
	encoding, err := json.Marshal([]string{})
	if err != nil {
		return err
	}
	// TODO: warn if file already exists
	return os.WriteFile(path, encoding, 0644)
}

type DatabaseHandler struct {
	path string
}

func GetDBHandler(dbpath string) DatabaseHandler {
	return DatabaseHandler{path: dbpath}
}

func (db *DatabaseHandler) ReadTodos() (Todos, error) {
	contents, err := os.ReadFile(db.path)
	if err != nil {
		return Todos{}, err
	}

	var todos Todos
	err = json.Unmarshal(contents, &todos)
	if err != nil {
		return Todos{}, err
	}
	return todos, nil
}

func (db *DatabaseHandler) WriteTodos(todoList Todos) error {
	encoding, err := json.Marshal(todoList)
	if err != nil {
		return err
	}
	return os.WriteFile(db.path, encoding, 0644)
}
