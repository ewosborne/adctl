/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "List all services AGH knows about",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listAll called")
	},
}

func init() {
	serviceCmd.AddCommand(listAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
