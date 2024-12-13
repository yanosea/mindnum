package command

import (
	"os"

	"github.com/yanosea/mindnum/app/presentation/cli/mindnum/formatter"
	"github.com/yanosea/mindnum/app/presentation/cli/mindnum/presenter"

	"github.com/yanosea/mindnum/pkg/proxy"
)

type Cli struct {
	Cobra       proxy.Cobra
	Version     string
	RootCommand proxy.Command
}

var (
	output string
	format = "text"
)

func NewCli(
	cobra proxy.Cobra,
	version string,
) *Cli {
	rootCommand := NewRootCommand(cobra, version, &format, &output)
	return &Cli{
		Cobra:       cobra,
		Version:     version,
		RootCommand: rootCommand,
	}
}

func (c *Cli) Run() int {
	out := os.Stdout
	exitCode := 0

	if err := c.RootCommand.Execute(); err != nil {
		output = formatter.AppendErrorToOutput(err, output)
		out = os.Stderr
		exitCode = 1
	}

	presenter.Present(out, output)

	return exitCode
}
