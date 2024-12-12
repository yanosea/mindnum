package mindnum

import (
	c "github.com/spf13/cobra"

	mindnumApp "github.com/yanosea/mindnum/application/mindnum"
	"github.com/yanosea/mindnum/presentation/cli/formatter"

	"github.com/yanosea/mindnum-pkg/proxy"
)

func NewVersionCommand(
	cobra proxy.Cobra,
	version string,
	format *string,
	output *string,
) proxy.Command {
	cmd := cobra.NewCommand()
	cmd.SetSilenceErrors(true)
	cmd.SetUse("version")
	cmd.SetHelpTemplate(versionHelpTemplate)
	cmd.SetArgs(cobra.ExactArgs(0))
	cmd.SetRunE(
		func(_ *c.Command, _ []string) error {
			return runVersion(version, format, output)
		},
	)

	return cmd
}

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
	versionHelpTemplate = `🔖 Show the version of mindnum

Usage:
  mindnum version [flags]

Flags:
  -h, --help  🤝 help for mindnum version
`
)
