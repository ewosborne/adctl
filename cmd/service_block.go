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

}
