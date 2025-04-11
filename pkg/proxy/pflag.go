package proxy

import (
	"github.com/spf13/pflag"
)

// FlagSet is an interface that provides a proxy of the methods of pflag.FlagSet.
type FlagSet interface {
	StringVarP(p *string, name string, shorthand string, value string, usage string)
}

// flagSetProxy is a proxy struct that implements the FlagSet interface.
type flagSetProxy struct {
	FlagSet *pflag.FlagSet
}

// StringVarP is a proxy method that calls the StringVarP method of the pflag.FlagSet.
func (f *flagSetProxy) StringVarP(p *string, name string, shorthand string, value string, usage string) {
	f.FlagSet.StringVarP(p, name, shorthand, value, usage)
}
