package main


/*
---------------------------------------------------------
Cobra can generate scripts that plug into your shell 
(bash, zsh, etc.) and enable autocomplete.

1. RegisterFlagCompletionFunc
2. MarkFlagFilename -> suggests matching file types
---------------------------------------------------------
*/

import (
    "fmt"

    "github.com/spf13/cobra"
)

var format string

func main() {
    rootCmd := &cobra.Command{
        Use:   "app",
        Short: "Demo CLI",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Format selected:", format)
        },
    }

    // Add flag
    rootCmd.Flags().StringVar(&format, "format", "", "Output format")

    // Add custom completion
    rootCmd.RegisterFlagCompletionFunc("format",
        func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

            return []string{"json", "yaml", "csv"}, cobra.ShellCompDirectiveDefault
        })

    rootCmd.Execute()
}