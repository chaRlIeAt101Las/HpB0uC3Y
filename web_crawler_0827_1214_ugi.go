// 代码生成时间: 2025-08-27 12:14:01
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"

    "github.com/PuerkitoBio/goquery"
    "github.com/spf13/cobra"
)

// mainCmd represents the base command when called without any subcommands
var mainCmd = &cobra.Command{
    Use:   "web_crawler",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
web_crawler [flags]
	This will fetch the URL provided in the -url flag and print its content.`,
    Run: func(cmd *cobra.Command, args []string) {
        runCrawler()
    },
}

// urlFlag holds the url value provided by the user
var urlFlag string

func init() {
    // Define a flag for the URL to fetch
    mainCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "URL to fetch content from")
}

// runCrawler fetches the content from the provided URL and prints it
func runCrawler() {
    if urlFlag == "" {
        fmt.Println("Error: URL is required")
        return
    }

    resp, err := http.Get(urlFlag)
    if err != nil {
        fmt.Printf("Error fetching URL: %v
", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %v
", err)
        return
    }

    // Use goquery to parse the HTML content
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        fmt.Printf("Error parsing HTML: %v
", err)
        return
    }

    // Print the content of the body tag or any other element you need to extract
    fmt.Println(doc.Find("body").Text())
}

// main entry point of the application
func main() {
    if err := mainCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
