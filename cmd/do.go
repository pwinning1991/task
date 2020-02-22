package cmd

import (
	"fmt"
	"strconv"

	"github.com/pwinning1991/task/db"
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

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number")
				continue
			}
			task := task[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}

		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
