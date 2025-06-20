package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

   var rootCmd = &cobra.Command{
       Use:   "todo",
       Short: "A CLI To-Do List application",
       Long:  `A simple CLI application to manage your to-do tasks.`,
   }

   func Execute() {
       if err := rootCmd.Execute(); err != nil {
           fmt.Fprintf(os.Stderr, "Error: %v\n", err)
           os.Exit(1)
       }
   }