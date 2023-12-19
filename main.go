package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// The main command that will be the root of all your subcommands
var rootCmd = &cobra.Command{
	Use:   "worklog",
	Short: "worklog lets you keep a log of your work",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists files",
	Run: func(cmd *cobra.Command, args []string) {
		ListFiles()
	},
}

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Writes files",
	Run: func(cmd *cobra.Command, args []string) {
		WriteFiles()
	},
}

func ListFiles() {
	fmt.Println("ListFiles function called")
}

func WriteFiles() {
	fmt.Println("WriteFiles function called")
}

func init() {
	// Add listCmd as a subcommand to rootCmd
	rootCmd.AddCommand(listCmd)
	// Add writeCmd as a subcommand to rootCmd
	rootCmd.AddCommand(writeCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
