package cmd

import (
	"log"
	"sort"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current tasks",
	Long: `Display the current set of tasks with their identifier,
    description, priority, and status.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		db := todo.GetDBHandler(dbPath)
		todos, err := db.ReadTodos()
		if err != nil {
			log.Fatalf("Failed to read todos from database with `%v`", err)
		}
		// Use Reverse to list higher priority first
		sort.Sort(sort.Reverse(todos))
		todos.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
