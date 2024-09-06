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

func (t *Todos) Add(todo_item Task) {
	todo_item.CreatedAt = time.Now()
	*t = append(*t, todo_item)
}

func (t *Todos) Complete(id int) error {
	todo_list := *t
	if id <= 0 || id > len(todo_list) {
		return errors.New("Index out of range")
	}

	// The Todo ids use 1-based indexing, so need to subtract 1
	todo_list[id-1].CompletedAt = time.Now()
	todo_list[id-1].Done = true

	return nil
}

func (t *Todos) Remove(id int) error {
	todo_list := *t
	if id <= 0 || id > len(todo_list) {
		return errors.New("Index out of range")
	}
	*t = append(todo_list[:id-1], todo_list[id:]...)
	return nil
}

func (t *Todos) Print() {
	for i, item := range *t {
		i++
		done_str := "false"
		if item.Done {
			done_str = "true"
		}
		fmt.Printf("%d - %s - Priority: %d - Done: %s\n", i, item.Description, item.Priority, done_str)
	}
}
