package cmd

import (
	"fmt"
	"os"
	"time"
	"todo-cli/tasks"

	"github.com/spf13/cobra"
)

   var completed bool
   var pending bool

   var listCmd = &cobra.Command{
       Use:   "list",
       Short: "List tasks",
       Run: func(cmd *cobra.Command, args []string) {
           taskList, err := tasks.LoadFromFile("todos.json")
           if err != nil {
               fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
               os.Exit(1)
           }
           showCompleted := completed
           showPending := pending
           if !showCompleted && !showPending {
               showCompleted, showPending = true, true
           }
           tasks := taskList.ListTasks(showCompleted, showPending)
           if len(tasks) == 0 {
               fmt.Println("No tasks found")
               return
           }
           for _, task := range tasks {
               status := "pending"
               if task.Completed {
                   status = "completed"
               }
               fmt.Printf("ID: %d, Description: %s, Status: %s, Priority: %d, Created: %s\n",
                   task.ID, task.Description, status, task.Priority, task.CreatedAt.Format(time.RFC822))
           }
       },
   }

   func init() {
       listCmd.Flags().BoolVar(&completed, "completed", false, "Show completed tasks")
       listCmd.Flags().BoolVar(&pending, "pending", false, "Show pending tasks")
       rootCmd.AddCommand(listCmd)
   }