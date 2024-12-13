package command

import (
	"testing"

	"github.com/yanosea/mindnum/pkg/proxy"
)

func TestNewRootCommand(t *testing.T) {
	type args struct {
		cobra   proxy.Cobra
		version string
		format  *string
		output  *string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive case",
			args: args{
				cobra:   proxy.NewCobra(),
				version: "v0.0.0",
				format:  nil,
				output:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRootCommand(tt.args.cobra, tt.args.version, tt.args.format, tt.args.output)
			if got.Execute() != nil {
				t.Errorf("NewRootCommand().Execute() returns an error")
			}
		})
	}
}
