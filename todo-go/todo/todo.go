package todo

import (
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

func (t *Todos) Print() {
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s - Priority: %d\n", i, item.Description, item.Priority)
	}
}
