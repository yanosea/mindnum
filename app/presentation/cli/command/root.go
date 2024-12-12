package command

import (
	"github.com/yanosea/mindnum/presentation/cli/command/mindnum"

	"github.com/yanosea/mindnum-pkg/proxy"
)

func NewRootCommand(
	cobra proxy.Cobra,
	version string,
	format *string,
	output *string,
) proxy.Command {
	cmd := cobra.NewCommand()
	cmd.SetSilenceErrors(true)
	cmd.SetUse("mindnum")
	cmd.SetHelpTemplate(rootHelpTemplate)
	cmd.SetVersion(version)
	cmd.AddCommand(
		mindnum.NewVersionCommand(
			cobra,
			version,
			format,
			output,
		),
		mindnum.NewGetCommand(
			cobra,
			format,
			output,
		),
	)

	return cmd
}

const (
	rootHelpTemplate = `🧠 mindnum is a CLI tool to get the mind number from the birthday

Usage:
  mindnum [command]

Available Commands:
  get         🧠 Get a mind number from an argument or the birthday flag and show the personality description
  help        🤝 Help about any command
  completion  🔧 Generate the autocompletion script for the specified shell
  version     🔖 Show the version of mindnum

Flags:
  -h, --help     🤝 help for mindnum
  -v, --version  🔖 version for mindnum

Use "mindnum [command] --help" for more information about a command.
`
)
