package cmd

import (
	"log"
	"strconv"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark a task as done",
	Long:  `Set the status of the given task to done`,
	Args:  cobra.MinimumNArgs(1),
	Run:   doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
	doneCmd.PersistentFlags().BoolP("verbose", "v", false,
		"Verbosity, whether to print task list after adding")
}

func doneRun(cmd *cobra.Command, args []string) {
	db := todo.GetDBHandler(dbPath)
	todos, err := db.ReadTodos()
	if err != nil {
		log.Fatalf("Failed to read todos from database with `%v`", err)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid ID\n", err)
	}
	todos.Complete(id)
	db.WriteTodos(todos)

	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		todos.Print()
	}
}
