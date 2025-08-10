// 代码生成时间: 2025-08-11 07:44:33
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "image"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/nfnt/resize"
    "github.com/spf13/cobra"
)

// Define the configuration structure for image resizing
type Config struct {
    Width  int
    Height int
    Target string
}

// Define the root command
var rootCmd = &cobra.Command{
    Use:   "resize_images",
    Short: "A utility to resize multiple images",
    Long:  "Resizes images to a specified width and height in a directory",
    Run:   run,
}

// Define the configuration variable with default values
var config Config

func init() {
    rootCmd.Flags().IntVarP(&config.Width, "width", "w", 1280, "Width of the resized image")
    rootCmd.Flags().IntVarP(&config.Height, "height", "h", 720, "Height of the resized image")
    rootCmd.Flags().StringVarP(&config.Target, "target", "t", "./images", "Directory containing images to resize")
}

// run is the main function that executes when the command is run
func run(cmd *cobra.Command, args []string) {
    // Check if the directory exists
    if _, err := os.Stat(config.Target); os.IsNotExist(err) {
        fmt.Println("Error: Target directory does not exist")
        return
    }

    // Process each file in the directory
    err := filepath.Walk(config.Target, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
