package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sprinkelcell/CLI-Task-Manager/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete your task",

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {

			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("'%s' this id is not correct\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.GetTaskList()
		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("This id '%d' is incorrect", id)
				continue
			}
			err := db.RemoveTask(tasks[id-1].Key)
			if err != nil {
				fmt.Printf("Something went wrong. Please try after some time")
				os.Exit(1)
			}
			fmt.Printf("id no '%d' marked as completed", id)
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)

}
