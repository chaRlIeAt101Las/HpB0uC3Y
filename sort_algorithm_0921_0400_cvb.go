// 代码生成时间: 2025-09-21 04:00:43
package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"

    "github.com/spf13/cobra"
)

// SortAlgorithmCmd represents the base command when called without any subcommands
var SortAlgorithmCmd = &cobra.Command{
    Use:   "sort-algorithm [algorithm] [list of numbers]",
    Short: "Sort a list of numbers using a specified algorithm",
    Long: `A brief description that
exposes the application functionality.`,
    Args:  cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        algorithm := args[0] // Get the sorting algorithm name
        numbersStr := args[1:]      // Get the list of numbers as strings
        numbers, err := strListToInt(numbersStr)
        if err != nil {
            fmt.Printf("Error converting numbers: %s
", err)
            os.Exit(1)
        }

        switch algorithm {
        case "bubble":
            bubbleSort(numbers)
        case "insertion":
            insertionSort(numbers)
        case "selection":
            selectionSort(numbers)
        default:
            fmt.Printf("Unknown algorithm: %s
", algorithm)
            os.Exit(1)
        }
    },
}

// strListToInt converts a slice of strings to a slice of integers
func strListToInt(strList []string) ([]int, error) {
    var intList []int
    for _, str := range strList {
        n, err := strconv.Atoi(str)
        if err != nil {
            return nil, fmt.Errorf("error converting string to int: %s", err)
        }
        intList = append(intList, n)
    }
    return intList, nil
}

// bubbleSort sorts a slice of integers using the bubble sort algorithm
func bubbleSort(nums []int) {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] {
                // Swap the elements
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
    fmt.Println(nums)
}

// insertionSort sorts a slice of integers using the insertion sort algorithm
func insertionSort(nums []int) {
    for i := 1; i < len(nums); i++ {
        key := nums[i]
        j := i - 1
        // Move elements of nums[0..i-1], that are greater than key, to one position ahead
        for j >= 0 && nums[j] > key {
            nums[j+1] = nums[j]
            j = j - 1
        }
        nums[j+1] = key
    }
    fmt.Println(nums)
}

// selectionSort sorts a slice of integers using the selection sort algorithm
func selectionSort(nums []int) {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        // Find the minimum element in the unsorted array
        minIdx := i
        for j := i + 1; j < n; j++ {
            if nums[j] < nums[minIdx] {
                minIdx = j
            }
        }
        // Swap the found minimum element with the first element
        nums[i], nums[minIdx] = nums[minIdx], nums[i]
    }
    fmt.Println(nums)
}

func main() {
    // Add the command to the root of the application
    rootCmd := &cobra.Command{Use: "sort-algorithm"}
    rootCmd.AddCommand(SortAlgorithmCmd)

    // Execute the root command
    err := rootCmd.Execute()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}