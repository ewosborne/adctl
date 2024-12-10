/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblock one or more services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unblock called")
	},
}

func init() {
	serviceCmd.AddCommand(unblockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unblockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unblockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
