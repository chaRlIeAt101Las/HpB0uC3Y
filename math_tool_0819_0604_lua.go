// 代码生成时间: 2025-08-19 06:04:55
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "os"

    "github.com/spf13/cobra"
)

// calcCmd represents the math command
var calcCmd = &cobra.Command{
    Use:   "math", // Brief name of command
    Short: "Mathematical operations", // Short description of command
    Long:  `A simple math tool set that can perform basic mathematical operations`, // Long description of command
    Run:   rootCmd,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "math-tool",
    Short: "Mathematical operations tool",
    Long:  `A tool to perform basic mathematical operations.`,
}

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add", // Brief name of command
    Short: "Add two numbers", // Short description of command
    Long:  `Add two numbers together.`, // Long description of command
    Run:   addNumbers,
}

// subCmd represents the subtract command
var subCmd = &cobra.Command{
    Use:   "subtract", // Brief name of command
    Short: "Subtract two numbers", // Short description of command
    Long:  `Subtract the second number from the first.`, // Long description of command
    Run:   subtractNumbers,
}

// mulCmd represents the multiply command
var mulCmd = &cobra.Command{
    Use:   "multiply", // Brief name of command
    Short: "Multiply two numbers", // Short description of command
    Long:  `Multiply two numbers together.`, // Long description of command
    Run:   multiplyNumbers,
}

// divCmd represents the divide command
var divCmd = &cobra.Command{
    Use:   "divide", // Brief name of command
    Short: "Divide two numbers", // Short description of command
    Long:  `Divide the first number by the second.`, // Long description of command
    Run:   divideNumbers,
}

// generateRandomCmd represents the generate-random command
var generateRandomCmd = &cobra.Command{
    Use:   "generate-random", // Brief name of command
    Short: "Generate a random number", // Short description of command
    Long:  `Generate a random number between 0 and 100.`, // Long description of command
    Run:   generateRandomNumber,
}

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add", // Brief name of command
    Short: "Add two numbers", // Short description of command
    Long:  `Add two numbers together.`, // Long description of command
    Run:   addNumbers,
}

func main() {
    // Initialize the random number generator
    rand.Seed(time.Now().UnixNano())

    // Add subcommands to the root command
    rootCmd.AddCommand(calcCmd)
    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(subCmd)
    rootCmd.AddCommand(mulCmd)
    rootCmd.AddCommand(divCmd)
    rootCmd.AddCommand(generateRandomCmd)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// addNumbers adds two numbers together
func addNumbers(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
        fmt.Println("You must provide two numbers to add")
        return
    }

    arg1, err1 := strconv.Atoi(args[0])
    arg2, err2 := strconv.Atoi(args[1])

    if err1 != nil || err2 != nil {
        fmt.Println("Both arguments must be integers")
        return
    }

    fmt.Printf("The sum of %d and %d is %d
", arg1, arg2, arg1+arg2)
}

// subtractNumbers subtracts the second number from the first
func subtractNumbers(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
        fmt.Println("You must provide two numbers to subtract")
        return
    }

    arg1, err1 := strconv.Atoi(args[0])
    arg2, err2 := strconv.Atoi(args[1])

    if err1 != nil || err2 != nil {
        fmt.Println("Both arguments must be integers")
        return
    }

    fmt.Printf("The difference between %d and %d is %d
", arg1, arg2, arg1-arg2)
}

// multiplyNumbers multiplies two numbers together
func multiplyNumbers(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
        fmt.Println("You must provide two numbers to multiply")
        return
    }

    arg1, err1 := strconv.Atoi(args[0])
    arg2, err2 := strconv.Atoi(args[1])

    if err1 != nil || err2 != nil {
        fmt.Println("Both arguments must be integers")
        return
    }

    fmt.Printf("The product of %d and %d is %d
", arg1, arg2, arg1*arg2)
}

// divideNumbers divides the first number by the second
func divideNumbers(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
        fmt.Println("You must provide two numbers to divide")
        return
    }

    arg1, err1 := strconv.Atoi(args[0])
    arg2, err2 := strconv.Atoi(args[1])

    if err1 != nil || err2 != nil {
        fmt.Println("Both arguments must be integers")
        return
    }

    if arg2 == 0 {
        fmt.Println("You cannot divide by zero")
        return
    }

    fmt.Printf("The result of %d divided by %d is %f
", arg1, arg2, float64(arg1)/float64(arg2))
}

// generateRandomNumber generates a random number between 0 and 100
func generateRandomNumber(cmd *cobra.Command, args []string) {
    fmt.Printf("Random number: %d
", rand.Intn(101))
}