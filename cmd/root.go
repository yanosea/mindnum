/*
Copyright © 2023 yanosea <myanoshi0626@gmail.com>
*/
package cmd

import (
	"embed"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version = ""

var resources embed.FS
var resourcePath = "resources/"

var rootCmd = &cobra.Command{
	Use:     "mindnum",
	Version: getVersion(version),
	Short:   "'mindnum' is a CLI tool to get the mind number from the birthday.",
	Long: `'mindnum' is a CLI tool to get the mind number from the birthday.

You can get the mind number from the birthday (yyyyMMdd).
You can check the character from the mind number.`,
}

func Execute(embedded embed.FS) {
	resources = embedded
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}

func getVersion(v string) string {
	// if version is embedded, return it
	if v != "" {
		return v
	}

	if i, ok := debug.ReadBuildInfo(); !ok {
		return "unknown"
	} else {
		return i.Main.Version
	}
}
