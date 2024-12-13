package mindnum

import (
	"testing"

	"github.com/yanosea/mindnum/pkg/proxy"
)

func TestNewGetCommand(t *testing.T) {
	type args struct {
		cobra  proxy.Cobra
		format *string
		output *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				cobra:  proxy.NewCobra(),
				format: func() *string { s := "text"; return &s }(),
				output: new(string),
			},
			// get command without the birthday flag or argument should return an error
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetCommand(tt.args.cobra, tt.args.format, tt.args.output)
			if err := got.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("NewGetCommand().Execute() returns an error")
			}
		})
	}
}

func Test_runGet(t *testing.T) {
	type args struct {
		birthday *string
		format   *string
		output   *string
		args     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive case (using the argument)",
			args: args{
				birthday: new(string),
				format:   func() *string { s := "text"; return &s }(),
				output:   new(string),
				args:     []string{"19900719"},
			},
			wantErr: false,
		},
		{
			name: "positive case (using the birthday flag)",
			args: args{
				birthday: func() *string { s := "19900719"; return &s }(),
				format:   func() *string { s := "text"; return &s }(),
				output:   new(string),
				args:     nil,
			},
			wantErr: false,
		},
		{
			name: "negative case (uc.Run() returns an error)",
			args: args{
				birthday: new(string),
				format:   func() *string { s := "text"; return &s }(),
				output:   new(string),
				args:     nil,
			},
			wantErr: true,
		},
		{
			name: "negative case (formatter.NewFormatter() returns an error)",
			args: args{
				birthday: func() *string { s := "19900719"; return &s }(),
				format:   new(string),
				output:   new(string),
				args:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := runGet(tt.args.birthday, tt.args.format, tt.args.output, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("runGet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
