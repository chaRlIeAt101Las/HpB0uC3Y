// 代码生成时间: 2025-08-08 23:11:49
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID        string  "json:"id""
    Name      string  "json:"name""
    Price     float64 "json:"price""
    Quantity  int     "json:"quantity"
}

// Cart represents the shopping cart
type Cart struct {
    Items map[string]CartItem
}

// NewCart initializes a new shopping cart
func NewCart() *Cart {
    return &Cart{
        Items: make(map[string]CartItem),
    }
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(item CartItem) error {
    if _, exists := c.Items[item.ID]; exists {
        return fmt.Errorf("item with id %s already exists in the cart", item.ID)
    }
    c.Items[item.ID] = item
    return nil
}

// RemoveItem removes an item from the cart
func (c *Cart) RemoveItem(itemID string) error {
    if _, exists := c.Items[itemID]; !exists {
        return fmt.Errorf("item with id %s does not exist in the cart", itemID)
    }
    delete(c.Items, itemID)
    return nil
}

// ListItems lists all items in the cart
func (c *Cart) ListItems() []CartItem {
    var items []CartItem
    for _, item := range c.Items {
        items = append(items, item)
    }
    return items
}

// CalculateTotal calculates the total price of all items in the cart
func (c *Cart) CalculateTotal() float64 {
    total := 0.0
    for _, item := range c.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "shopping-cart", // program name
    Short: "A brief description of your application", // short description
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This app is a tool to generate the needed files
to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {
        cart := NewCart()
        // Add items to the cart
        if err := cart.AddItem(CartItem{ID: "1", Name: "Apple", Price: 0.99, Quantity: 3}); err != nil {
            log.Fatal(err)
        }
        if err := cart.AddItem(CartItem{ID: "2", Name: "Banana", Price: 1.49, Quantity: 2}); err != nil {
            log.Fatal(err)
        }

        // List items in the cart
        items := cart.ListItems()
        jsonData, err := json.MarshalIndent(items, "", "  ")
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(string(jsonData))

        // Calculate total price
        total := cart.CalculateTotal()
        fmt.Printf("Total price: %.2f
", total)
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
// It only needs to happen once to the RootCmd
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}