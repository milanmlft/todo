package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

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
	Run: listRun,
}

var doneOpt bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' tasks")
}

func listRun(cmd *cobra.Command, args []string) {
	db := todo.GetDBHandler(dbPath)
	todos, err := db.ReadTodos()
	if err != nil {
		log.Fatalf("Failed to read todos from database with `%v`", err)
	}
	// Use Reverse to list higher priority first
	sort.Sort(sort.Reverse(todos))

	writer := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, task := range todos {
		if task.Done == doneOpt {
			fmt.Fprintln(writer, task.PrettyPrint())
		}
	}
	writer.Flush()
}
