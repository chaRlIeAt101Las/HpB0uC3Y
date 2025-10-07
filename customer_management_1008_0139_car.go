// 代码生成时间: 2025-10-08 01:39:30
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/spf13/cobra"
)

// Customer represents a customer in the customer relationship management system.
type Customer struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"createdAt"`
}

// customerData holds the data of all customers.
var customerData = make(map[int]Customer)

// NewCustomer creates a new customer and adds it to the customer data map.
func NewCustomer(id int, name, email string) *Customer {
    customer := Customer{
        ID:        id,
        Name:      name,
        Email:     email,
        CreatedAt: time.Now(),
    }
    customerData[id] = customer
    return &customer
}

// GetCustomer retrieves a customer by their ID.
func GetCustomer(id int) (*Customer, error) {
    if customer, ok := customerData[id]; ok {
        return &customer, nil
    }
    return nil, fmt.Errorf("customer with ID %d not found", id)
}

// DeleteCustomer removes a customer by their ID from the customer data map.
func DeleteCustomer(id int) error {
    if _, ok := customerData[id]; ok {
        delete(customerData, id)
        return nil
    }
    return fmt.Errorf("customer with ID %d not found", id)
}

// CustomerCmd represents the customer command in the CLI.
var CustomerCmd = &cobra.Command{
    Use:   "customer",
    Short: "Manage customers",
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

// AddCustomerCmd represents the add customer command in the CLI.
var AddCustomerCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new customer",
    Args:  cobra.ExactArgs(3),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            log.Fatalf("Invalid ID: %s", err)
        }
        name := args[1]
        email := args[2]
        NewCustomer(id, name, email)
        fmt.Println("Customer added successfully")
    },
}

// GetCustomerCmd represents the get customer command in the CLI.
var GetCustomerCmd = &cobra.Command{
    Use:   "get",
    Short: "Get a customer by ID",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            log.Fatalf("Invalid ID: %s", err)
        }
        customer, err := GetCustomer(id)
        if err != nil {
            log.Fatalf("Error getting customer: %s", err)
        }
        data, _ := json.Marshal(customer)
        fmt.Println(string(data))
    },
}

// DeleteCustomerCmd represents the delete customer command in the CLI.
var DeleteCustomerCmd = &cobra.Command{
    Use:   "delete",
    Short: "Delete a customer by ID",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            log.Fatalf("Invalid ID: %s", err)
        }
        err = DeleteCustomer(id)
        if err != nil {
            log.Fatalf("Error deleting customer: %s", err)
        }
        fmt.Println("Customer deleted successfully")
    },
}

func init() {
    CustomerCmd.AddCommand(AddCustomerCmd)
    CustomerCmd.AddCommand(GetCustomerCmd)
    CustomerCmd.AddCommand(DeleteCustomerCmd)
}

func main() {
    if err := CustomerCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}