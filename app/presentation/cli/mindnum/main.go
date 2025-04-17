package main

import (
	"os"

	"github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum/command"

	"github.com/yanosea/mindnum/v2/pkg/proxy"
	"github.com/yanosea/mindnum/v2/pkg/utility"
)

var (
	// version is the version of mindnum cli.
	version = ""
	// cobra is a proxy of spf13/cobra.
	cobra = proxy.NewCobra()
	// versionUtil is a provider of the version of the application.
	versionUtil = utility.NewVersionUtil(proxy.NewDebug())
	// exit is a proxy of os.Exit.
	exit = os.Exit
)

// main is the entry point of mindnum cli.
func main() {
	cli := command.NewCli(
		cobra,
		versionUtil.GetVersion(version),
	)
	exit(cli.Run())
}
