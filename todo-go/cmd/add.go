package cmd

import (
	"log"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the todo list with the option to set a pririty level.`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().BoolP("verbose", "v", false,
		"Verbosity, whether to print task list after adding")
}

func addRun(cmd *cobra.Command, args []string) {
	db := todo.GetDBHandler(dbPath)
	todos, err := db.ReadTodos()
	if err != nil {
		log.Fatalf("Failed to read todos from database with `%v`", err)
	}

	for _, arg := range args {
		newTask := todo.Task{Description: arg}
		todos.Add(newTask)
	}
	db.WriteTodos(todos)

	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		todos.Print()
	}
}
