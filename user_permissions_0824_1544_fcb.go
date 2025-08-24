// 代码生成时间: 2025-08-24 15:44:31
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
)

// Permission defines the structure for a user permission
type Permission struct {
    UserID    int    `json:"user_id"`
    Permission string `json:"permission"`
}

// UserPermissionsManager is the main structure for managing user permissions
type UserPermissionsManager struct {
    permissions []Permission
}

// NewUserPermissionsManager creates a new instance of UserPermissionsManager
func NewUserPermissionsManager() *UserPermissionsManager {
    return &UserPermissionsManager{
        permissions: make([]Permission, 0),
    }
}

// AddPermission adds a new permission to the manager
func (m *UserPermissionsManager) AddPermission(userID int, permission string) error {
    m.permissions = append(m.permissions, Permission{UserID: userID, Permission: permission})
    return nil
}

// RemovePermission removes a permission from the manager
func (m *UserPermissionsManager) RemovePermission(userID int, permission string) error {
    for i, p := range m.permissions {
        if p.UserID == userID && p.Permission == permission {
            m.permissions = append(m.permissions[:i], m.permissions[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("permission not found for user ID: %d", userID)
}

// ListPermissions lists all permissions in the manager
func (m *UserPermissionsManager) ListPermissions() error {
    if len(m.permissions) == 0 {
        return fmt.Errorf("no permissions found")
    }
    for _, p := range m.permissions {
        fmt.Printf("User ID: %d, Permission: %s
", p.UserID, p.Permission)
    }
    return nil
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
    Use:   "user-permissions",
    Short: "Manage user permissions",
    Long:  `A brief description of your application`,
    Run: func(cmd *cobra.Command, args []string) {
        if err := RootCmd.Execute(); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    },
}

// rootCmd represents the most basic command getting called here,
// it's used to set the flags which are common to subcommands,
// and it's used to invoke the subcommands.
var rootCmd = &cobra.Command{
    Use:   "user-permissions",
    Short: "Manage user permissions",
    Long:  `This is a user permission management system using Cobra framework`,
}

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new user permission",
    Run: func(cmd *cobra.Command, args []string) {
        userManager := NewUserPermissionsManager()
        userID, _ := cmd.Flags().GetInt("user_id")
        permission, _ := cmd.Flags().GetString("permission")
        if err := userManager.AddPermission(userID, permission); err != nil {
            log.Fatalf("Error adding permission: %s", err)
        }
    },
}

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
    Use:   "remove",
    Short: "Remove a user permission",
    Run: func(cmd *cobra.Command, args []string) {
        userManager := NewUserPermissionsManager()
        userID, _ := cmd.Flags().GetInt("user_id")
        permission, _ := cmd.Flags().GetString("permission")
        if err := userManager.RemovePermission(userID, permission); err != nil {
            log.Fatalf("Error removing permission: %s", err)
        }
    },
}

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all user permissions",
    Run: func(cmd *cobra.Command, args []string) {
        userManager := NewUserPermissionsManager()
        if err := userManager.ListPermissions(); err != nil {
            log.Fatalf("Error listing permissions: %s", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(removeCmd)
    rootCmd.AddCommand(listCmd)

    // Here you will define your flags and configuration settings.
    addCmd.Flags().Int("user_id", 0, "User ID for permission addition")
    addCmd.Flags().String("permission", "", "Permission to add")
    removeCmd.Flags().Int("user_id", 0, "User ID for permission removal")
    removeCmd.Flags().String("permission", "", "Permission to remove")
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
