package cmd

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli/tasks"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit ID DESCRIPTION",
	Short: "Edit a task's description",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil || id <= 0 {
			fmt.Println("Error: Invalid ID")
			return
		}
		description := args[1]
		if description == "" {
			fmt.Println("Error: Description cannot be empty")
			return
		}
		taskList, err := tasks.LoadFromFile("todos.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}
		if err := taskList.EditTask(id, description); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		if err := taskList.SaveToFile("todos.json"); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d updated: %s\n", id, description)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}