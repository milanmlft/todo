package cmd

import (
	"log"
	"strconv"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Long: `Remove a task from the todo list based on its ID.
    It's possible to delete multiple tasks at once.`,
	Args: cobra.MinimumNArgs(1),
	Run:  runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().BoolP("verbose", "v", false,
		"Verbosity, whether to print the task list when command finishes")
}

func runRemove(cmd *cobra.Command, args []string) {
	db := todo.GetDBHandler(dbPath)
	todos, err := db.ReadTodos()
	if err != nil {
		log.Fatalf("Failed to read todos from database with `%v`", err)
	}

	for i, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln(args[0], "is not a valid ID\n", err)
		}

		// Need to shift ids in case of multiple args, as removing an item
		// will shift the position of subsequent items down
		truePosition := id - i
		err = todos.Remove(truePosition)
		if err != nil {
			log.Fatalf("Error for arg %s:  %v", arg, err)
		}
	}
	db.WriteTodos(todos)
	if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
		todos.Print()
	}
}
