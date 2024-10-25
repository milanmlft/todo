package cmd

import (
	"log"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the todo list with the option to set a pririty level.`,
	Run:   addRun,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2,
		"Set priority 1, 2 or 3")
}

func addRun(cmd *cobra.Command, args []string) {
	db := todo.GetDBHandler(viper.GetString("datafile"))
	todos, err := db.ReadTodos()
	if err != nil {
		log.Fatalf("Failed to read todos from database with `%v`", err)
	}

	for _, arg := range args {
		newTask := todo.Task{Description: arg}
		newTask.SetPriority(priority)
		todos.Add(newTask)
		log.Println("Added", newTask.Description)
	}
	db.WriteTodos(todos)
}
