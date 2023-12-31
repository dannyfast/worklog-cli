package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type FileDetails struct {
	Name string `json:"name"`
	Size string `json:"size"`
}

type Payload struct {
	Files []FileDetails `json:"files"`
	Total int           `json:"total"`
}

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

func getListOfFiles() ([]FileDetails, int) {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return nil, 0
	}

	var totalSize int64
	var files []FileDetails

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("Error getting info for file: %s, error: %v\n", entry.Name(), err)
				continue
			}

			size := info.Size()
			totalSize += size
			humanSize := fileSizeToHumanReadable(size)
			files = append(files, FileDetails{Name: entry.Name(), Size: humanSize})
		}
	}

	return files, int(totalSize)
}

func ListFiles() {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				fmt.Printf("Error getting info for file: %s, error: %v\n", entry.Name(), err)
				continue
			}

			fullPath, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error getting full path: %v\n", err)
				return
			}

			size := info.Size()
			humanSize := fileSizeToHumanReadable(size)
			fmt.Printf("%v/%v - %v\n", fullPath, entry.Name(), humanSize)
		}
	}
}

func fileSizeToHumanReadable(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func WriteFiles() {
	files, totalSize := getListOfFiles()
	payload := Payload{
		Files: files,
		Total: totalSize,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling payload: %v\n", err)
		return
	}

	// Replace with the actual API endpoint you want to call
	apiEndpoint := "http://localhost/api/write" //TODO: build this later

	resp, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Error calling API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("API Response: %s\n", string(body))
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
