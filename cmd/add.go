package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sprinkelcell/CLI-Task-Manager/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task in your task manger",

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Printf("Something went wrong. Please try after some time")
			os.Exit(1)
		}
		fmt.Printf("'%s' added successfully.", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
