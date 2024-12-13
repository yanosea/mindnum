package mindnum

import (
	c "github.com/spf13/cobra"

	mindnumApp "github.com/yanosea/mindnum/app/application/mindnum"
	mindnumRepo "github.com/yanosea/mindnum/app/infrastructure/text/repository"
	"github.com/yanosea/mindnum/app/presentation/cli/mindnum/formatter"

	"github.com/yanosea/mindnum/pkg/proxy"
)

var (
	birthday string
)

func NewGetCommand(
	cobra proxy.Cobra,
	format *string,
	output *string,
) proxy.Command {
	cmd := cobra.NewCommand()
	cmd.SetSilenceErrors(true)
	cmd.SetUse("get [birthday (optional)]")
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
	getHelpTemplate = `🧠 Get a mind number from an argument or the birthday flag and show the personality description

Usage:
  mindnum get [birthday (optional)] [flags]

Flags:
  -b, --birthday string  🎂 birthday in the format 'yyyyMMdd' (e.g. 19900719) (optional)
  -h, --help             🤝 help for mindnum get

Arguments:
  birthday  🎂 birthday in the format 'yyyyMMdd' (e.g. 19900719) (optional)
`
)
