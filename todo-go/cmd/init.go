package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/milanmlft/todo/todo-go/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a new todo list.",
	Long: `Create a new database to store your todo list.

    For example:
        todo-go init                        Creates new database at default location ~/.todo-go.json
        todo-go init --db-path todos.json   Creates new database at ./todos.json
    `,
	Args: cobra.MaximumNArgs(0),
	Run:  initRun,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initRun(cmd *cobra.Command, args []string) {
	// TODO: ask for confirmation to overwrite existing file
	dbpath := viper.GetString("datafile")
	if fileExists(dbpath) {
		log.Printf("Database already exists at %s. Doing nothing.", dbpath)
		return
	}
	err := todo.InitialiseDB(dbpath)
	if err != nil {
		log.Fatalf("Initialising database failed with error `%v`", err)
	}
	log.Printf("Initialised todo database at %s", viper.GetString("datafile"))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Printf("Error checking file %s: %v", path, err)
	return false
}
