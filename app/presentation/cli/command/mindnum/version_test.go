package mindnum

import (
	"fmt"
	"testing"

	"github.com/yanosea/mindnum-pkg/proxy"
)

func TestNewVersionCommand(t *testing.T) {
	type args struct {
		cobra   proxy.Cobra
		version string
		format  *string
		output  *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				cobra:   proxy.NewCobra(),
				version: "v0.0.0",
				format:  func() *string { s := "text"; return &s }(),
				output:  new(string),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVersionCommand(tt.args.cobra, tt.args.version, tt.args.format, tt.args.output)
			if err := got.Execute(); (err != nil) != tt.wantErr {
				fmt.Println(err)
				t.Errorf("NewVersionCommand().Execute() returns an error")
			}
		})
	}
}

func Test_runVersion(t *testing.T) {
	type args struct {
		version string
		format  *string
		output  *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				version: "v0.0.0",
				format:  func() *string { s := "text"; return &s }(),
				output:  new(string),
			},
			wantErr: false,
		},
		{
			name: "negative case (formatter.NewFormatter() returns an error)",
			args: args{
				version: "v0.0.0",
				format:  new(string),
				output:  new(string),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runVersion(tt.args.version, tt.args.format, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("runVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
