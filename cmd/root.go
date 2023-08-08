/*
Copyright © 2023 yanosea <myanoshi0626@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = ""

var rootCmd = &cobra.Command{
	Use:     "mindnum",
	Version: version,
	Short:   "'mindnum' is a CLI tool to get the mind number from the birthday.",
	Long: `'mindnum' is a CLI tool to get the mind number from the birthday.

You can get the mind number from the birthday (yyyyMMdd).`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
