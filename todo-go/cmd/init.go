package cmd

import (
	"log"

	"github.com/milanmlft/todo/todo-go/database"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		dbPath = "~/.todo-go.json"
	case 1:
		dbPath = args[0]
	default:
		log.Fatal("init requires at most 1 argument.")

	}
	err := database.InitialiseDB(dbPath)
	if err != nil {
		log.Fatalf("Initialising database failed with error `%v`", err)
	}
}
