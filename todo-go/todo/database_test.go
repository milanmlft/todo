package todo

import (
	"os"
	"testing"
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
	origTodos := Todos{Task{Description: "todo"}}

	err = db.WriteTodos(origTodos)
	if err != nil {
		t.Fatalf("db.WriteTodos(mockTodos) failed with %v; want nil", err)
	}

	readTodos, err := db.ReadTodos()
	if err != nil {
		t.Fatalf("db.ReadTodos() failed with %v; want nil", err)
	}
	if len(readTodos) != len(origTodos) {
		t.Fatalf("db.ReadTodos() returned %d todos, want %d", len(readTodos), len(origTodos))
	}
	if readTodos[0].Description != origTodos[0].Description {
		t.Fatalf("db.ReadTodos() returned non-matching description: %v, want %v",
			readTodos[0].Description, origTodos[0].Description)
	}
	if readTodos[0].Priority != origTodos[0].Priority {
		t.Fatalf("db.ReadTodos() returned non-matching priority: %d, want %d",
			readTodos[0].Priority, origTodos[0].Priority)
	}
	if readTodos[0].Done != origTodos[0].Done {
		t.Fatalf("db.ReadTodos() returned non-matching done status: %v, want %v",
			readTodos[0].Done, origTodos[0].Done)
	}
}

func TestReadingTodosPosition(t *testing.T) {
	dbpath, err := os.CreateTemp("", "tmp.json")
	if err != nil {
		panic(err)
	}
	defer os.Remove(dbpath.Name())

	db := DatabaseHandler{path: dbpath.Name()}
	origTodos := Todos{Task{Description: "todo"}}

	err = db.WriteTodos(origTodos)
	if err != nil {
		t.Fatalf("db.WriteTodos(mockTodos) failed with %v; want nil", err)
	}

	readTodos, err := db.ReadTodos()
	if err != nil {
		t.Fatalf("db.ReadTodos() failed with %v; want nil", err)
	}

	if readTodos[0].position != 1 {
		t.Fatalf("db.ReadTodos() did not set correct position: %v, want %v",
			readTodos[0].position, 1)
	}
}
