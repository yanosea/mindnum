package command

import (
	"errors"
	"testing"

	"github.com/yanosea/mindnum/pkg/proxy"

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
			name: "positive case",
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
	type fields struct {
		Cobra       proxy.Cobra
		Version     string
		RootCommand proxy.Command
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		setup  func(mockCtrl *gomock.Controller, tt *fields)
	}{
		{
			name: "positive case",
			fields: fields{
				Cobra:       proxy.NewCobra(),
				Version:     "v0.0.0",
				RootCommand: NewRootCommand(proxy.NewCobra(), "v0.0.0", &format, &output),
			},
			want:  0,
			setup: nil,
		},
		{
			name: "negative case",
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			if tt.setup != nil {
				tt.setup(mockCtrl, &tt.fields)
			}
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
