package main

import (
	"github.com/milanmlft/todo/todo-go/todo"
)

func main() {
	todos := &todo.Todos{}
	task := todo.Task{
		Description: "Do this",
		Priority:    2,
	}
	todos.Add(task)
	todos.Print()
}
