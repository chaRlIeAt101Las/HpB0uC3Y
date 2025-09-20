// 代码生成时间: 2025-09-20 20:02:10
// db_pool_manager.go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "github.com/spf13/cobra"
)

// DBConfig defines the configuration for database connection.
type DBConfig struct {
    Username string
    Password string
    Host     string
    Port     string
    DBName  string
}

// Pool represents a database connection pool.
type Pool struct {
    *sql.DB
}

// NewPool creates and returns a new database connection pool.
func NewPool(cfg DBConfig) (*Pool, error) {
    // DSN (Data Source Name) for MySQL database connection.
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

    // Open the database connection.
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool.
    db.SetMaxIdleConns(10)

    // Set the maximum lifetime of a connection.
    db.SetConnMaxLifetime(5 * time.Minute)

    // Ping the database to verify the connection.
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Return the new Pool.
    return &Pool{db}, nil
}

// Close closes the database, releasing any open resources.
func (p *Pool) Close() error {
    return p.DB.Close()
}

func main() {
    var dbConfig DBConfig
    var rootCmd = &cobra.Command{
        Use:   "dbpool",
        Short: "Manages database connection pools",
        Run: func(cmd *cobra.Command, args []string) {
            dbPool, err := NewPool(dbConfig)
            if err != nil {
                log.Fatalf("Failed to create DB pool: %v", err)
            }
            defer dbPool.Close()

            fmt.Println("Database connection pool created and ready to use.")
        },
    }

    rootCmd.Flags().StringVar(&dbConfig.Username, "username", "your_username", "Database username")
    rootCmd.Flags().StringVar(&dbConfig.Password, "password", "your_password", "Database password")
    rootCmd.Flags().StringVar(&dbConfig.Host, "host", "localhost", "Database host")
    rootCmd.Flags().StringVar(&dbConfig.Port, "port", "3306", "Database port")
    rootCmd.Flags().StringVar(&dbConfig.DBName, "dbname", "your_db_name", "Database name")

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Error executing the root command: %v", err)
    }
}
