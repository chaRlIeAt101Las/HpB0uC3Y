// 代码生成时间: 2025-09-19 11:01:02
package main

import (
    "bufio"
    "bytes"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

// ErrorLogger is the main application structure
type ErrorLogger struct {
    verbose bool
    logFile string
}

// NewErrorLogger creates a new instance of ErrorLogger
func NewErrorLogger(verbose bool, logFile string) *ErrorLogger {
    return &ErrorLogger{
        verbose: verbose,
        logFile: logFile,
    }
}

// Run starts the error logger
func (e *ErrorLogger) Run() error {
    // Open the log file
    file, err := os.OpenFile(e.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open log file: %w", err)
    }
    defer file.Close()

    // Create a writer that writes to both file and stdout if verbose
    var writer *bufio.Writer
    if e.verbose {
        writer = bufio.NewWriterSize(&os.StdoutAndErrWriter{}, 1024)
    } else {
        writer = bufio.NewWriter(file)
    }
    defer writer.Flush()

    // Read from stdin and write to log
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if e.verbose {
            fmt.Println(line)
        }
        _, err := writer.WriteString(line + "
")
        if err != nil {
            return fmt.Errorf("failed to write to log: %w", err)
        }
    }
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("failed to read input: %w", err)
    }
    return nil
}

// NewCmdErrorLogger creates a new cobra command for ErrorLogger
func NewCmdErrorLogger() *cobra.Command {
    verbose := false
    logFile := "error.log"

    cmd := &cobra.Command{
        Use:   "error-logger",
        Short: "Collects error logs",
        Long:  `ErrorLogger is a command-line tool to collect error logs.`,
        Run: func(cmd *cobra.Command, args []string) {
            el := NewErrorLogger(verbose, logFile)
            if err := el.Run(); err != nil {
                log.Fatal(err)
            }
        },
    }

    cmd.Flags().BoolVarP(&verbose, "verbose", "v", verbose, "Enable verbose output")
    cmd.Flags().StringVarP(&logFile, "log-file", "l", logFile, "Specify the log file path")

    return cmd
}

func main() {
    rootCmd := NewCmdErrorLogger()
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}