package command

import (
	"errors"
	"io"
	"testing"

	"github.com/yanosea/mindnum/v2/app/presentation/cli/mindnum/presenter"
	"github.com/yanosea/mindnum/v2/pkg/proxy"

	"go.uber.org/mock/gomock"
)

func TestNewCli(t *testing.T) {
	type args struct {
		cobra   proxy.Cobra
		version string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive testing",
			args: args{
				cobra:   proxy.NewCobra(),
				version: "v0.0.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCli(tt.args.cobra, tt.args.version)
			if got.Run() != 0 {
				t.Errorf("NewCli().Run() returns an error")
			}
		})
	}
}

func TestCli_Run(t *testing.T) {
	origPrint := presenter.Print

	type fields struct {
		Cobra       proxy.Cobra
		Version     string
		RootCommand proxy.Command
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		setup   func(mockCtrl *gomock.Controller, tt *fields)
		cleanup func()
	}{
		{
			name: "positive testing",
			fields: fields{
				Cobra:       proxy.NewCobra(),
				Version:     "v0.0.0",
				RootCommand: NewRootCommand(proxy.NewCobra(), "v0.0.0", &format, &output),
			},
			want:  0,
			setup: nil,
		},
		{
			name: "negative testing (c.RootCommand.Execute() failed)",
			fields: fields{
				Cobra:       proxy.NewCobra(),
				Version:     "v0.0.0",
				RootCommand: NewRootCommand(proxy.NewCobra(), "v0.0.0", &format, &output),
			},
			want: 1,
			setup: func(mockCtrl *gomock.Controller, tt *fields) {
				mockCommand := proxy.NewMockCommand(mockCtrl)
				mockCommand.EXPECT().Execute().Return(errors.New("test"))
				tt.RootCommand = mockCommand
			},
			cleanup: nil,
		},
		{
			name: "negative testing (presenter.Print failed)",
			fields: fields{
				Cobra:       proxy.NewCobra(),
				Version:     "v0.0.0",
				RootCommand: NewRootCommand(proxy.NewCobra(), "v0.0.0", &format, &output),
			},
			want: 1,
			setup: func(mockCtrl *gomock.Controller, tt *fields) {
				presenter.Print = func(writer io.Writer, output string) error {
					return errors.New("print error")
				}
			},
			cleanup: func() {
				presenter.Print = origPrint
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			if tt.setup != nil {
				tt.setup(mockCtrl, &tt.fields)
			}
			defer func() {
				if tt.cleanup != nil {
					tt.cleanup()
				}
			}()
			c := &Cli{
				Cobra:       tt.fields.Cobra,
				Version:     tt.fields.Version,
				RootCommand: tt.fields.RootCommand,
			}
			if got := c.Run(); got != tt.want {
				t.Errorf("Cli.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
