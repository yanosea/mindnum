package proxy

import (
	"io"

	"github.com/spf13/cobra"
)

// Cobra is an interface that provides a proxy of the methods of cobra
type Cobra interface {
	ExactArgs(int) PositionalArgs
	MaximumNArgs(int) PositionalArgs
	NewCommand() Command
}

// cobraProxy is a proxy struct that implements the Cobra interface.
type cobraProxy struct{}

// NewCobra returns a new instance of the Cobra interface.
func NewCobra() Cobra {
	return &cobraProxy{}
}

// ExactArgs is a proxy method that calls the ExactArgs method of the cobra.ExactArgs.
func (*cobraProxy) ExactArgs(n int) PositionalArgs {
	return &positionalArgsProxy{PositionalArgs: cobra.ExactArgs(n)}
}

// MaximumNArgs is a proxy method that calls the MaximumNArgs method of the cobra.MaximumNArgs.
func (*cobraProxy) MaximumNArgs(n int) PositionalArgs {
	return &positionalArgsProxy{PositionalArgs: cobra.MaximumNArgs(n)}
}

// NewCommand returns a new instance of the Command interface.
func (*cobraProxy) NewCommand() Command {
	return &commandProxy{Command: &cobra.Command{}}
}

// PositionalArgs is an interface that provides a proxy of the methods of cobra.PositionalArgs.
type PositionalArgs interface {
	GetPositionalArgs() cobra.PositionalArgs
}

// positionalArgsProxy is a proxy struct that implements the PositionalArgs interface.
type positionalArgsProxy struct {
	PositionalArgs cobra.PositionalArgs
}

// GetPositionalArgs is a proxy method that calls the GetPositionalArgs method of the cobra.PositionalArgs.
func (p *positionalArgsProxy) GetPositionalArgs() cobra.PositionalArgs {
	return p.PositionalArgs
}

// Command is an interface that provides a proxy of the methods of cobra.Command.
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

// commandProxy is a proxy struct that implements the Command interface.
type commandProxy struct {
	Command *cobra.Command
}

// AddCommand is a proxy method that calls the AddCommand method of the cobra.Command.
func (c *commandProxy) AddCommand(cmds ...Command) {
	for _, cmd := range cmds {
		c.Command.AddCommand(cmd.GetCommand())
	}
}

// Execute is a proxy method that calls the Execute method of the cobra.Command.
func (c *commandProxy) Execute() error {
	return c.Command.Execute()
}

// GetCommand is a proxy method that calls the GetCommand method of the cobra.Command.
func (c *commandProxy) GetCommand() *cobra.Command {
	return c.Command
}

// PersistentFlags is a proxy method that calls the PersistentFlags method of the cobra.Command.
func (c *commandProxy) PersistentFlags() FlagSet {
	return &flagSetProxy{FlagSet: c.Command.PersistentFlags()}
}

// RunE is a proxy method that calls the RunE method of the cobra.Command.
func (c *commandProxy) RunE(cmd *cobra.Command, args []string) error {
	return c.Command.RunE(cmd, args)
}

// SetArgs is a proxy method that calls the SetArgs method of the cobra.Command.
func (c *commandProxy) SetArgs(positionalArgs PositionalArgs) {
	c.Command.Args = positionalArgs.GetPositionalArgs()
}

// SetErr is a proxy method that calls the SetErr method of the cobra.Command.
func (c *commandProxy) SetErr(io io.Writer) {
	c.Command.SetErr(io)
}

// SetHelpTemplate is a proxy method that calls the SetHelpTemplate method of the cobra.Command.
func (c *commandProxy) SetHelpTemplate(s string) {
	c.Command.SetHelpTemplate(s)
}

// SetMaximumNArgs is a proxy method that calls the SetMaximumNArgs method of the cobra.Command.
func (c *commandProxy) SetMaximumNArgs(n int) {
	c.Command.Args = cobra.MaximumNArgs(n)
}

// SetOut is a proxy method that calls the SetOut method of the cobra.Command.
func (c *commandProxy) SetOut(io io.Writer) {
	c.Command.SetOut(io)
}

// SetRunE is a proxy method that calls the SetRunE method of the cobra.Command.
func (c *commandProxy) SetRunE(runE func(cmd *cobra.Command, args []string) error) {
	c.Command.RunE = runE
}

// SetSilenceErrors is a proxy method that calls the SetSilenceErrors method of the cobra.Command.
func (c *commandProxy) SetSilenceErrors(silenceErrors bool) {
	c.Command.SilenceErrors = silenceErrors
}

// SetUse is a proxy method that calls the SetUse method of the cobra.Command.
func (c *commandProxy) SetUse(use string) {
	c.Command.Use = use
}

// SetUsageTemplate is a proxy method that calls the SetUsageTemplate method of the cobra.Command.
func (c *commandProxy) SetUsageTemplate(s string) {
	c.Command.SetUsageTemplate(s)
}

// SetVersion is a proxy method that calls the SetVersion method of the cobra.Command.
func (c *commandProxy) SetVersion(version string) {
	c.Command.Version = version
}
