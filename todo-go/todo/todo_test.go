package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	todos := Todos{}
	task := Task{
		Description: "Do this",
		Priority:    2,
	}
	todos.Add(task)
	if len(todos) != 1 {
		t.Fatalf("todos.Add(task) gives len(todos) = %d; want 1", len(todos))
	}
}

func TestCompleteExisting(t *testing.T) {
	todos := Todos{}
	task := Task{
		Description: "Do this",
		Priority:    2,
	}
	todos.Add(task)
	err := todos.Complete(1)
	if err != nil {
		t.Fatalf("todos.Complete(1) = %v; want nil", err)
	}
}

func TestCompleteEmpty(t *testing.T) {
	todos := Todos{}
	err := todos.Complete(0)
	if err == nil {
		t.Fatalf("todos.Complete(1) on empty todos returned nil, want 'Index out of range'")
	}
}

func TestRemoveExisting(t *testing.T) {
	todos := Todos{}
	task := Task{
		Description: "Do this",
		Priority:    2,
	}
	todos.Add(task)
	err := todos.Remove(1)
	remainingTodos := len(todos)
	if remainingTodos != 0 || err != nil {
		t.Fatalf("todos.Remove(1) = %v and len(todos) = %d; want nil and len(todos) = 0", err, remainingTodos)
	}
}

func TestRemoveEmpty(t *testing.T) {
	todos := Todos{}
	err := todos.Remove(1)
	if err == nil {
		t.Fatalf("todos.Remove(1) on empty todos returned nil, want 'Index out of range'")
	}
}
