package todo

import "testing"

func generateTodos(n int) Todos {
	var todoList Todos
	for i := 0; i < n; i++ {
		todoList = append(todoList, Task{})
	}
	return todoList
}

func TestLen(t *testing.T) {
	maxLen := 100
	for i := 0; i < maxLen; i++ {
		todos := generateTodos(i)
		if todos.Len() != i {
			t.Fatalf("todos.Len() returned %d, want %d", todos.Len(), i)
		}
	}
}

func TestSwap(t *testing.T) {
	todos := Todos{Task{Description: "a"}, Task{Description: "b"}}
	todos.Swap(0, 1)
	if todos[0].Description != "b" || todos[1].Description != "a" {
		t.Fatalf("{todos[0].Description, todos[1].Description} = {%s, %s}, want {%s, %s}",
			todos[0].Description, todos[1].Description, "b", "a")
	}
}

func TestLessSamePriority(t *testing.T) {
	todos := Todos{Task{Priority: 2, position: 1}, Task{Priority: 2, position: 2}}
	if !todos.Less(0, 1) {
		t.Fatalf("todos.Less(0, 1) is false, want true")
	}
}

func TestLessDifferentPriority(t *testing.T) {
	todos := Todos{Task{Priority: 3, position: 1}, Task{Priority: 2, position: 2}}
	if !todos.Less(1, 0) {
		t.Fatalf("todos.Less(1, 0) is false, want true")
	}
}
