package cmd

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli/tasks"

	"github.com/spf13/cobra"
)

var priorityCmd = &cobra.Command{
	Use:   "priority ID PRIORITY",
	Short: "Set task priority",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil || id <= 0 {
			fmt.Println("Error: Invalid ID")
			return
		}
		priority, err := strconv.Atoi(args[1])
		if err != nil || priority < 1 || priority > 3 {
			fmt.Println("Error: Priority must be 1 (high), 2 (medium), or 3 (low)")
			return
		}
		taskList, err := tasks.LoadFromFile("todos.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}
		if err := taskList.SetPriority(id, priority); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		if err := taskList.SaveToFile("todos.json"); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d priority set to %d\n", id, priority)
	},
}

func init() {
	rootCmd.AddCommand(priorityCmd)
}
