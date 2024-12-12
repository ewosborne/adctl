/*
Copyright © 2024 Eric Osborne
No header.
*/
package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var debugLogger *log.Logger
var outputFormat string
var enableDebug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "adctl",
	Version: "0.5.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&enableDebug, "debug", "d", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output format", "o", "json", "Enable debug mode")

	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)

	// need PreRun because flags aren't parsed until a command is run.
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if enableDebug {
			debugLogger.SetOutput(os.Stderr)
		} else {
			debugLogger.SetOutput(io.Discard)
		}
	}
}
