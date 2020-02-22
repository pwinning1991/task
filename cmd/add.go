package cmd

import (
	"fmt"
	"strings"

	"github.com/pwinning1991/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
