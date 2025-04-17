package command

import (
	"os"

	"github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum/formatter"
	"github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum/presenter"

	"github.com/yanosea/mindnum/v2/pkg/proxy"
)

// Cli is a struct that represents the command line interface of mindnum cli.
type Cli struct {
	Cobra       proxy.Cobra
	Version     string
	RootCommand proxy.Command
}

var (
	// output is the output string
	output string
	// format is the format of the output
	format = "text"
)

// NewCli returns a new instance of the Cli struct.
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

// Run runs the command line interface of mindnum cli.
func (c *Cli) Run() int {
	out := os.Stdout
	exitCode := 0

	if err := c.RootCommand.Execute(); err != nil {
		output = formatter.AppendErrorToOutput(err, output)
		out = os.Stderr
		exitCode = 1
	}

	if err := presenter.Print(out, output); err != nil {
		exitCode = 1
	}

	return exitCode
}
