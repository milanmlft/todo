package cmd

import (
	"log"
	"strconv"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark tasks as done",
	Long:  `Set the status of the given tasks to done`,
	Args:  cobra.MinimumNArgs(1),
	Run:   doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func doneRun(cmd *cobra.Command, args []string) {
	db := todo.GetDBHandler(viper.GetString("datafile"))
	todos, err := db.ReadTodos()
	if err != nil {
		log.Fatalf("Failed to read todos from database with `%v`", err)
	}

	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln(arg, "is not a valid ID\n", err)
		}
		err = todos.Complete(id)
		if err != nil {
			log.Fatalf("%d did not match any items", id)
		}
		log.Println(id, "marked done")
	}
	db.WriteTodos(todos)
}
