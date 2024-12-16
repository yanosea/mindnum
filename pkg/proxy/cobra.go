package proxy

import (
	"io"

	"github.com/spf13/cobra"
)

type Cobra interface {
	ExactArgs(int) PositionalArgs
	MaximumNArgs(int) PositionalArgs
	NewCommand() Command
}

type cobraProxy struct{}

func NewCobra() Cobra {
	return &cobraProxy{}
}

func (*cobraProxy) ExactArgs(n int) PositionalArgs {
	return &positionalArgsProxy{PositionalArgs: cobra.ExactArgs(n)}
}

func (*cobraProxy) MaximumNArgs(n int) PositionalArgs {
	return &positionalArgsProxy{PositionalArgs: cobra.MaximumNArgs(n)}
}

func (*cobraProxy) NewCommand() Command {
	return &commandProxy{Command: &cobra.Command{}}
}

type PositionalArgs interface {
	GetPositionalArgs() cobra.PositionalArgs
}

type positionalArgsProxy struct {
	PositionalArgs cobra.PositionalArgs
}

func (p *positionalArgsProxy) GetPositionalArgs() cobra.PositionalArgs {
	return p.PositionalArgs
}

type Command interface {
	AddCommand(cmds ...Command)
	Execute() error
	GetCommand() *cobra.Command
	PersistentFlags() FlagSet
	RunE(cmd *cobra.Command, args []string) error
	SetArgs(positionalArgs PositionalArgs)
	SetErr(io io.Writer)
	SetHelpTemplate(s string)
	SetMaximumNArgs(n int)
	SetOut(io io.Writer)
	SetRunE(runE func(cmd *cobra.Command, args []string) error)
	SetSilenceErrors(silenceErrors bool)
	SetUse(use string)
	SetUsageTemplate(s string)
	SetVersion(version string)
}

type commandProxy struct {
	Command *cobra.Command
}

func (c *commandProxy) AddCommand(cmds ...Command) {
	for _, cmd := range cmds {
		c.Command.AddCommand(cmd.GetCommand())
	}
}

func (c *commandProxy) Execute() error {
	return c.Command.Execute()
}

func (c *commandProxy) GetCommand() *cobra.Command {
	return c.Command
}

func (c *commandProxy) PersistentFlags() FlagSet {
	return &flagSetProxy{FlagSet: c.Command.PersistentFlags()}
}

func (c *commandProxy) RunE(cmd *cobra.Command, args []string) error {
	return c.Command.RunE(cmd, args)
}

func (c *commandProxy) SetArgs(positionalArgs PositionalArgs) {
	c.Command.Args = positionalArgs.GetPositionalArgs()
}

func (c *commandProxy) SetErr(io io.Writer) {
	c.Command.SetErr(io)
}

func (c *commandProxy) SetHelpTemplate(s string) {
	c.Command.SetHelpTemplate(s)
}

func (c *commandProxy) SetMaximumNArgs(n int) {
	c.Command.Args = cobra.MaximumNArgs(n)
}

func (c *commandProxy) SetOut(io io.Writer) {
	c.Command.SetOut(io)
}

func (c *commandProxy) SetRunE(runE func(cmd *cobra.Command, args []string) error) {
	c.Command.RunE = runE
}

func (c *commandProxy) SetSilenceErrors(silenceErrors bool) {
	c.Command.SilenceErrors = silenceErrors
}

func (c *commandProxy) SetUse(use string) {
	c.Command.Use = use
}

func (c *commandProxy) SetUsageTemplate(s string) {
	c.Command.SetUsageTemplate(s)
}

func (c *commandProxy) SetVersion(version string) {
	c.Command.Version = version
}
