package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run as a client.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo: impl client")
	},
}
