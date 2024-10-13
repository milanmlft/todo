package main

import (
	"log"

	"github.com/milanmlft/todo/todo-go/todo"
)

func main() {
	todos := &todo.Todos{}
	task := todo.Task{
		Description: "Do this",
		Priority:    2,
	}
	log.Println("Adding first todo")
	todos.Add(task)
	todos.Print()

	log.Println("Marking first todo as complete")
	todos.Complete(1)
	todos.Print()

	log.Println("Adding second todo")
	todos.Add(todo.Task{Description: "Do that"})
	todos.Print()
	log.Println("Deleting second todo")
	todos.Remove(2)
	todos.Print()
}
