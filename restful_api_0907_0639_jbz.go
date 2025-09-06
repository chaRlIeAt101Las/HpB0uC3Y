// 代码生成时间: 2025-09-07 06:39:18
package main

import (
	"fmt"
	"net/http"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Define a rootCmd as a cobra.Command
var rootCmd = &cobra.Command{
	Use:   "restful-api",
	Short: "A simple RESTful API server",
	Long:  `A simple RESTful API server built with Cobra and Go.`,
}

// main function to run the server
func main() {
	// Add routes to the server
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", rootHandler)

	// Start the server
	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// rootHandler handles the root path
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the RESTful API server!
")
}

// helloHandler handles the /hello path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world!
")
}

// init is called before main and is used to initialize the Cobra command
func init() {
	// Set the default command for Cobra if no arguments are passed
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
