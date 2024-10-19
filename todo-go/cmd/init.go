package cmd

import (
	"log"
	"os"

	"github.com/milanmlft/todo/todo-go/database"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [optional path]",
	Short: "Initialise a new todo list.",
	Long: `Create a new database to store your todo list.

    For example:
        todo-go init                Creates new database at default location ~/.todo-go.json
        todo-go init todos.json     Creates new database at ./todso.json
    `,
	Run: initRun,
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initRun(cmd *cobra.Command, args []string) {
	var dbPath string
	switch len(args) {
	case 0:
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to retrieve home directory. Set database path manually.")
		}
		dbPath = home + string(os.PathSeparator) + ".todo-go.json"
	case 1:
		dbPath = args[0]
	default:
		log.Fatal("init requires at most 1 argument.")

	}
	err := database.InitialiseDB(dbPath)
	if err != nil {
		log.Fatalf("Initialising database failed with error `%v`", err)
	}
	log.Printf("Initialised todo database at %s", dbPath)
}
