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
	Long: "Add, delete, or list DNS rewrites.",
}

func init() {
	rootCmd.AddCommand(rewriteCmd)
}
