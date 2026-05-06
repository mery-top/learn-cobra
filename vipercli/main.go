package main
/*
---------------------------------------------------------

Viper: Used for automatic configuration Management

Used for fetching data from Multiple resources

---------------------------------------------------------
*/


import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
    Use:   "app",
    Short: "Demo app",
    Run: func(cmd *cobra.Command, args []string) {
        port := viper.GetInt("port")
        debug := viper.GetBool("debug")

        fmt.Println("Port:", port)
        fmt.Println("Debug:", debug)
    },
}

func init() {
    // Defaults
    viper.SetDefault("port", 8080)
    viper.SetDefault("debug", false)

    // Config file setup
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig()
    if err == nil {
        fmt.Println("Using config:", viper.ConfigFileUsed())
    }

    // Environment variables
    viper.AutomaticEnv()

    // Flags
    rootCmd.Flags().Int("port", 3000, "Port number")
    rootCmd.Flags().Bool("debug", false, "Enable debug")

    // Bind flags to Viper
    viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
    viper.BindPFlag("debug", rootCmd.Flags().Lookup("debug"))
}

func main() {
    rootCmd.Execute()
}