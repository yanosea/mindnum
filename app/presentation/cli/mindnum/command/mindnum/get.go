package mindnum

import (
	c "github.com/spf13/cobra"

	mindnumApp "github.com/yanosea/mindnum/app/application/mindnum"
	mindnumRepo "github.com/yanosea/mindnum/app/infrastructure/text/repository"
	"github.com/yanosea/mindnum/app/presentation/cli/mindnum/formatter"

	"github.com/yanosea/mindnum/pkg/proxy"
)

var (
	// birthday is the receiver of the birthday flag
	birthday string
)

// NewGetCommand returns a new instance of the get command.
func NewGetCommand(
	cobra proxy.Cobra,
	format *string,
	output *string,
) proxy.Command {
	cmd := cobra.NewCommand()
	cmd.SetSilenceErrors(true)
	cmd.SetUse("get [birthday (optional)]")
	cmd.SetAliases([]string{"gt", "g"})
	cmd.SetUsageTemplate(getUsageTemplate)
	cmd.SetHelpTemplate(getHelpTemplate)
	cmd.PersistentFlags().StringVarP(
		&birthday,
		"birthday",
		"b",
		"",
		"Birthday in the format 'yyyyMMdd' (e.g. 20000101) (optional)",
	)
	cmd.SetRunE(
		func(_ *c.Command, args []string) error {
			return runGet(&birthday, format, output, args)
		},
	)

	return cmd
}

// runGet runs the get command.
func runGet(birthday *string, format *string, output *string, args []string) error {
	if len(args) > 0 {
		*birthday = args[0]
	}

	mindnumRepo := mindnumRepo.NewMindnumRepository()
	uc := mindnumApp.NewGetMindnumUseCase(mindnumRepo)
	dto, err := uc.Run(*birthday)
	if err != nil {
		return err
	}

	f, err := formatter.NewFormatter(*format)
	if err != nil {
		return err
	}

	o := f.Format(dto)

	*output = o

	return nil
}

const (
	// getHelpTemplate is the help template of the get command.
	getHelpTemplate = `🧠 Get a mind number from an argument or the birthday flag and show the personality description

` + getUsageTemplate
	// getUsageTemplate is the usage template of the get command.
	getUsageTemplate = `Usage:
  mindnum get [flags] [argument]

Flags:
  -b, --birthday string  🎂 birthday in the format 'yyyyMMdd' (e.g. 19900719) (optional)
  -h, --help             🤝 help for mindnum get

Arguments:
  birthday  🎂 birthday in the format 'yyyyMMdd' (e.g. 19900719) (optional)
`
)
