package stringer

/*
---------------------------------------------------------
Sub cmd: stringer rev/reverse arg[0]

---------------------------------------------------------
*/
import (
    "fmt"

    "stringer/pkg/stringer"
    "github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
    Use:   "reverse",
    Aliases: []string{"rev"},
    Short:  "Reverses a string",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        res := stringer.Reverse(args[0])
        fmt.Println(res)
    },
}

func init() {
    rootCmd.AddCommand(reverseCmd)
}