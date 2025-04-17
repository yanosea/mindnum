package mindnum

import (
	c "github.com/spf13/cobra"

	mindnumApp "github.com/yanosea/mindnum/v2/app/application/mindnum"
	"github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum/formatter"

	"github.com/yanosea/mindnum/v2/pkg/proxy"
)

// NewVersionCommand returns a new instance of the version command.
func NewVersionCommand(
	cobra proxy.Cobra,
	version string,
	format *string,
	output *string,
) proxy.Command {
	cmd := cobra.NewCommand()
	cmd.SetSilenceErrors(true)
	cmd.SetUse("version")
	cmd.SetUsageTemplate(versionUsageTemplate)
	cmd.SetHelpTemplate(versionHelpTemplate)
	cmd.SetArgs(cobra.ExactArgs(0))
	cmd.SetRunE(
		func(_ *c.Command, _ []string) error {
			return runVersion(version, format, output)
		},
	)

	return cmd
}

// runVersion runs the version command.
func runVersion(version string, format *string, output *string) error {
	uc := mindnumApp.NewVersionUseCase()
	dto := uc.Run(version)

	f, err := formatter.NewFormatter(*format)
	if err != nil {
		return err
	}

	o := f.Format(dto)

	*output = o

	return nil
}

const (
	// versionHelpTemplate is the help template of the version command.
	versionHelpTemplate = `🔖 Show the version of mindnum

` + versionUsageTemplate
	// versionUsageTemplate is the usage template of the version command.
	versionUsageTemplate = `Usage:
  mindnum version [flags]

Flags:
  -h, --help  🤝 help for mindnum version
`
)
