/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rewriteCmd represents the rewrite command
var rewriteCmd = &cobra.Command{
	Use:   "rewrite",
	Short: "Control DNS rewrites",
	RunE:  RewriteCmdE,
}

func RewriteCmdE(cmd *cobra.Command, args []string) error {
	// doesn't do much by itself.

	return nil
}

func init() {
	rootCmd.AddCommand(rewriteCmd)
}
