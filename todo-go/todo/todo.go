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
	position    int
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

func (task *Task) PrettyPrint() string {
	var priorityLabel string

	switch task.Priority {
	case 1:
		priorityLabel = "low"
	case 3:
		priorityLabel = "high"
	default:
		priorityLabel = "medium"
	}

	doneLabel := "[ ]"
	if task.Done {
		doneLabel = "[x]"
	}

	return fmt.Sprintf(
		"%d. \t %s \t %s \t Priority: %s",
		task.position, doneLabel, task.Description, priorityLabel)
}
