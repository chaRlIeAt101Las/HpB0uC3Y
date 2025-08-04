// 代码生成时间: 2025-08-04 22:29:25
 * This program creates a CLI application with a command to insert data into a database,
 * ensuring SQL injection prevention by using prepared statements.
 */

package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"
    "github.com/go-sql-driver/mysql"
    "github.com/spf13/cobra"
)

// Define custom error type
type AppError struct {
    Err error
}

// Database configurations
const (
    dbUser     = "username"
    dbPassword = "password"
    dbHost     = "127.0.0.1"
    dbName     = "database"
)

// Connection string
const connString = "%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"

// Initialize database connection
var db *sql.DB
var err error

func initDB() {
    db, err = sql.Open("mysql", fmt.Sprintf(connString, dbUser, dbPassword, dbHost, dbName))
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()
}

// Execute a query with input parameters to prevent SQL injection
func executeQuery(query string, args ...interface{}) (*sql.Rows, error) {
    rows, err := db.Query(query, args...)
    if err != nil {
        return nil, AppError{Err: err}
    }
    return rows, nil
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "anti-sql-injection",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines
and likely contains examples to demonstrate how to use the application.`,
}

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "A brief description of what the command does",
    Long: `A longer description of what the command does.`,
    Run: func(cmd *cobra.Command, args []string) {
        // A pointer to a struct that will be used to insert data into the database
        data := struct {
            Name     string
            Email    string
            Created time.Time
        }{}
        
        fmt.Print("Enter name: ")
        fmt.Scanln(&data.Name)
        
        fmt.Print("Enter email: ")
        fmt.Scanln(&data.Email)
        
        fmt.Print("Enter created date (YYYY-MM-DD): ")
        fmt.Scanln(&data.Created)
        
        // Prepare the query to prevent SQL injection
        query := `INSERT INTO users (name, email, created) VALUES (?, ?, ?)`
        
        // Execute the prepared query with input parameters
        _, err := executeQuery(query, data.Name, data.Email, data.Created)
        if err != nil {
            fmt.Println("Failed to add data: ", err)
            return
        }
        fmt.Println("Data added successfully.")
    },
}

func main() {
    initDB()
    
    RootCmd.AddCommand(addCmd)
    
    if err := RootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}