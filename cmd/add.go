package cmd

import (
	"fmt"
	"os"
	"todo-cli/tasks"

	"github.com/spf13/cobra"
)

var priority int

var addCmd = &cobra.Command{
    Use:   "add DESCRIPTION",
    Short: "Add a new task",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        description := args[0]
        taskList, err := tasks.LoadFromFile("todos.json")
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
            os.Exit(1)
        }
        if err := taskList.AddTask(description, priority); err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            return
        }
        if err := taskList.SaveToFile("todos.json"); err != nil {
            fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
            os.Exit(1)
        }
        fmt.Printf("Task added: %s (Priority: %d)\n", description, priority)
    },
}

func init() {
    addCmd.Flags().IntVarP(&priority, "priority", "p", 3, "Priority of the task (1=high, 2=medium, 3=low)")
    rootCmd.AddCommand(addCmd)
}