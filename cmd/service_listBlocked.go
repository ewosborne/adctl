/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listBlockedCmd represents the listBlocked command
var listBlockedCmd = &cobra.Command{
	Use:   "listBlocked",
	Short: "List all blocked services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listBlocked called")
	},
}

func init() {
	serviceCmd.AddCommand(listBlockedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listBlockedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listBlockedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
