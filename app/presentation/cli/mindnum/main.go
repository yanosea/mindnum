package main

import (
	"os"

	"github.com/yanosea/mindnum/app/presentation/cli/mindnum/command"

	"github.com/yanosea/mindnum/pkg/proxy"
	"github.com/yanosea/mindnum/pkg/utility"
)

var (
	version     = ""
	cobra       = proxy.NewCobra()
	versionUtil = utility.NewVersionUtil(proxy.NewDebug())
	exit        = os.Exit
)

func main() {
	cli := command.NewCli(
		cobra,
		versionUtil.GetVersion(version),
	)
	exit(cli.Run())
}
