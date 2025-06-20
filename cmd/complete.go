package cmd

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli/tasks"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete ID",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil || id <= 0 {
			fmt.Println("Error: Invalid ID")
			return
		}
		taskList, err := tasks.LoadFromFile("todos.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}
		if taskList.CompleteTask(id) {
			fmt.Printf("Task %d marked as completed\n", id)
		} else {
			fmt.Printf("Task %d not found\n", id)
		}
		if err := taskList.SaveToFile("todos.json"); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
