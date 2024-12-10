/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "List, block, and unblock services",
}

func init() {
	rootCmd.AddCommand(serviceCmd)

}
