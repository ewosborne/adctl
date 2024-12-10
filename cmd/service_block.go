/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Block one or more services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("block called")
	},
}

func init() {
	serviceCmd.AddCommand(blockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
