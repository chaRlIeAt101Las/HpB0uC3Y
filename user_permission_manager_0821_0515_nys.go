// 代码生成时间: 2025-08-21 05:15:56
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/spf13/cobra"
)

// Permission represents a user permission with an ID and Name
type Permission struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

// UserManager holds the permissions for a user
type UserManager struct {
    permissions map[uint]Permission
}

// NewUserManager creates a new instance of UserManager
func NewUserManager() *UserManager {
    return &UserManager{
        permissions: make(map[uint]Permission),
    }
}

// AddPermission adds a new permission to the user
func (um *UserManager) AddPermission(id uint, name string) error {
    if _, exists := um.permissions[id]; exists {
        return fmt.Errorf("permission with id %d already exists", id)
    }
    um.permissions[id] = Permission{id, name}
    return nil
}

// RemovePermission removes a permission from the user
func (um *UserManager) RemovePermission(id uint) error {
    if _, exists := um.permissions[id]; !exists {
        return fmt.Errorf("permission with id %d does not exist", id)
    }
    delete(um.permissions, id)
    return nil
}

// ListPermissions lists all permissions
func (um *UserManager) ListPermissions() []Permission {
    var permissions []Permission
    for _, perm := range um.permissions {
        permissions = append(permissions, perm)
    }
    return permissions
}

// userCmd represents the user permission management command
var userCmd = &cobra.Command{
    Use:   "user",
    Short: "Manage user permissions",
    Long:  `This command allows you to manage user permissions.`,
    Run: func(cmd *cobra.Command, args []string) {
        manager := NewUserManager()
        // Example usage: Add and list permissions
        if err := manager.AddPermission(1, "read"); err != nil {
            log.Fatalf("Error adding permission: %s", err)
        }
        if err := manager.AddPermission(2, "write"); err != nil {
            log.Fatalf("Error adding permission: %s", err)
        }
        perms := manager.ListPermissions()
        b, _ := json.MarshalIndent(perms, "", "  ")
        fmt.Println(string(b))
    },
}

func main() {
    // Set up signal handler for graceful shutdown
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        log.Println("Received signal, shutting down...")
        time.Sleep(3 * time.Second)
        os.Exit(0)
    }()

    // Add user permission management command
    userCmd.AddCommand(&cobra.Command{
        Use:   "add [id] [name]",
        Short: "Add a new permission",
        Args:  cobra.ExactArgs(2),
        Run: func(cmd *cobra.Command, args []string) {
            id, err := strconv.Atoi(args[0])
            if err != nil {
                log.Fatalf("Invalid ID: %s", err)
            }
            name := args[1]
            manager := NewUserManager()
            if err := manager.AddPermission(uint(id), name); err != nil {
                log.Fatalf("Error adding permission: %s", err)
            }
        },
    })

    userCmd.AddCommand(&cobra.Command{
        Use:   "remove [id]",
        Short: "Remove a permission",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            id, err := strconv.Atoi(args[0])
            if err != nil {
                log.Fatalf("Invalid ID: %s", err)
            }
            manager := NewUserManager()
            if err := manager.RemovePermission(uint(id)); err != nil {
                log.Fatalf("Error removing permission: %s", err)
            }
        },
    })

    userCmd.AddCommand(&cobra.Command{
        Use:   "list",
        Short: "List all permissions",
        Run: func(cmd *cobra.Command, args []string) {
            manager := NewUserManager()
            perms := manager.ListPermissions()
            b, _ := json.MarshalIndent(perms, "", "  ")
            fmt.Println(string(b))
        },
    })

    if err := userCmd.Execute(); err != nil {
        log.Fatalf("Error executing command: %s", err)
    }
}
