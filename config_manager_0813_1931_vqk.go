// 代码生成时间: 2025-08-13 19:31:16
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

// config represents the application configuration
type config struct {
    FileName string `default:"config.yaml"` // default configuration file name
}

// rootCmd is the main command for the application
var rootCmd = &cobra.Command{
    Use:   "config-manager",
    Short: "Manage configuration files",
    Long:  `A command line application to manage configuration files.`,
    Run: func(cmd *cobra.Command, args []string) {
        cmd.Help()
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func main() {
    Execute()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile, err := cmd.Flags().GetString("config"); err == nil { // config file flag is provided
        if _, err := os.Stat(cfgFile); !os.IsNotExist(err) {
            viper.SetConfigFile(cfgFile)
        } else {
            fmt.Printf("Config file %s not found, ignoring.", cfgFile)
        }
    }
    
    viper.AutomaticEnv() // read in environment variables that match
}

// addCommands adds subcommands to the root command
func addCommands() {
    // Define a command to load a configuration file
    var loadCmd = &cobra.Command{
        Use:   "load",
        Short: "Load a configuration file",
        Long:  `Loads a configuration file into the application.`,
        Run: func(cmd *cobra.Command, args []string) {
            if err := loadConfigFile(); err != nil {
                log.Fatalf("Error loading config: %v", err)
            }
            fmt.Println("Configuration loaded successfully.")
        },
    }
    rootCmd.AddCommand(loadCmd)

    // Define a command to view the current configuration
    var viewCmd = &cobra.Command{
        Use:   "view",
        Short: "View the current configuration",
        Long:  `Displays the current configuration settings.`,
        Run: func(cmd *cobra.Command, args []string) {
            if err := viewConfig(); err != nil {
                log.Fatalf("Error viewing config: %v", err)
            }
        },
    }
    rootCmd.AddCommand(viewCmd)
}

// loadConfigFile loads a configuration file from the specified path
func loadConfigFile() error {
    cfgFile, err := rootCmd.Flags().GetString("config")
    if err != nil {
        return err
    }
    
    viper.SetConfigName(filepath.Base(cfgFile)) // name of config file (without extension)
    viper.AddConfigPath(filepath.Dir(cfgFile)) // path to look for the config file in
    viper.ReadInConfig()
    return nil
}

// viewConfig prints out the current configuration settings
func viewConfig() error {
    currentConfig := viper.GetViper().AllSettings()
    fmt.Printf("Current Configuration: %+v", currentConfig)
    return nil
}

func init() {
    cobra.OnInitialize(initConfig)
    rootCmd.PersistentFlags().String("config", "", "config file (default is $HOME/.config/appname.yaml)")
    addCommands()
}
