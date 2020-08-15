package cmd

import (
	"fmt"
	"os"

	"github.com/sprinkelcell/CLI-Task-Manager/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of task in task manger",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetTaskList()
		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Printf("You have no task to complete")
			return
		}
		fmt.Println("You have following task to complete")
		for index, task := range tasks {
			fmt.Printf("%d. %s\n", index+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
