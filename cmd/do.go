package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks tasks completed",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}

		}

		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
