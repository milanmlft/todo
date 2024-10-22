package todo

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	CreatedAt   time.Time
	CompletedAt time.Time
	Description string
	Done        bool
	Priority    int
}

type Todos []Task

func (t *Todos) Add(todoItem Task) {
	todoItem.CreatedAt = time.Now()
	*t = append(*t, todoItem)
}

func (t *Todos) Complete(id int) error {
	todoList := *t
	if id <= 0 || id > len(todoList) {
		return errors.New("Index out of range")
	}

	// The Todo ids use 1-based indexing, so need to subtract 1
	todoList[id-1].CompletedAt = time.Now()
	todoList[id-1].Done = true

	return nil
}

func (t *Todos) Remove(id int) error {
	todoList := *t
	if id <= 0 || id > len(todoList) {
		return errors.New("Index out of range")
	}
	*t = append(todoList[:id-1], todoList[id:]...)
	return nil
}

func (t *Todos) Print() {
	for i, item := range *t {
		i++
		done := "false"
		if item.Done {
			done = "true"
		}
		fmt.Printf("%d - %s - Priority: %d - Done: %s\n", i, item.Description, item.Priority, done)
	}
}

func (task *Task) SetPriority(priority int) {
	// We only allow a fixed set of priorities between 1-3
	switch priority {
	case 1:
		task.Priority = 1
	case 3:
		task.Priority = 3
	default:
		task.Priority = 2
	}
}
