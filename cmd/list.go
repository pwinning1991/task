package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Used to list tasks left to do",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
