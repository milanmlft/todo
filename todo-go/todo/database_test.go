package database

import (
	"os"
	"testing"

	"github.com/milanmlft/todo/todo-go/todo"
)

func TestInitialiseDB(t *testing.T) {
	path, err := os.CreateTemp("", "tmp.json")
	if err != nil {
		panic(err)
	}
	defer os.Remove(path.Name())

	InitialiseDB(path.Name())
	result, err := os.ReadFile(path.Name())
	if err != nil {
		t.Fatalf("Failed to read from %s; %v", path.Name(), err)
	}
	resultString := string(result)
	want := "[]"
	if resultString != want {
		t.Fatalf("InitialiseDB() wrote %s; want %s", resultString, want)
	}
}

func TestInitialiseDBErrors(t *testing.T) {
	err := InitialiseDB("~/fail.json")
	if err == nil {
		t.Fatalf("InitialiseDB(\"~/fail.json\") gave nil error; want non-nill error")
	}
}

func TestReadingAndWritingTodos(t *testing.T) {
	dbpath, err := os.CreateTemp("", "tmp.json")
	if err != nil {
		panic(err)
	}
	defer os.Remove(dbpath.Name())

	db := DatabaseHandler{path: dbpath.Name()}
	origTodos := todo.Todos{todo.Task{Description: "todo"}}

	err = db.WriteTodos(origTodos)
	if err != nil {
		t.Fatalf("db.WriteTodos(mockTodos) failed with %v; want nil", err)
	}

	readTodos, err := db.ReadTodos()
	if len(readTodos) != len(origTodos) {
		t.Fatalf("db.readTodos() returned %d todos, want %d", len(readTodos), len(origTodos))
	}
	if readTodos[0] != origTodos[0] {
		t.Fatalf("db.ReadTodos() returned %v, want %v", readTodos, origTodos)
	}
}
