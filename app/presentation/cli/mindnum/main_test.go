package main

import (
	"os"
	"testing"

	mindnumRepo "github.com/yanosea/mindnum/v2/app/infrastructure/text/repository"

	"github.com/yanosea/mindnum/v2/pkg/proxy"
	"github.com/yanosea/mindnum/v2/pkg/utility"
)

func Test_main(t *testing.T) {
	capturer := utility.NewCapturer(
		proxy.NewOs(),
		proxy.NewBuffer(),
		proxy.NewBuffer(),
	)

	origExit := exit
	exit = func(code int) {}
	defer func() {
		exit = origExit
	}()

	type fields struct {
		fnc      func()
		capturer utility.Capturer
	}
	tests := []struct {
		name       string
		fields     fields
		args       []string
		wantStdOut string
		wantStdErr string
		wantErr    bool
	}{
		{
			name: "positive testing (root)",
			fields: fields{
				fnc: func() {
					main()
				},
				capturer: capturer,
			},

			args: []string{"path/to/mindnum"},
			wantStdOut: `🧠 mindnum is a CLI tool to get the mind number from the birthday

Usage:
  mindnum [command]

Available Commands:
  get         🧠 Get a mind number from an argument or the birthday flag and show the personality description
  help        🤝 Help about any command
  completion  🔧 Generate the autocompletion script for the specified shell
  version     🔖 Show the version of mindnum

Flags:
  -h, --help     🤝 help for mindnum
  -v, --version  🔖 version for mindnum

Use "mindnum [command] --help" for more information about a command.

`,
			wantStdErr: "",
			wantErr:    false,
		},
		{
			name: "positive testing (version)",
			fields: fields{
				fnc: func() {
					origVersion := "v0.0.0"
					version = origVersion
					defer func() {
						version = origVersion
					}()
					main()
				},
				capturer: capturer,
			},

			args:       []string{"path/to/mindnum", "version"},
			wantStdOut: "mindnum version v0.0.0\n",
			wantStdErr: "",
			wantErr:    false,
		},
		{
			name: "positive testing (get)",
			fields: fields{
				fnc: func() {
					main()
				},
				capturer: capturer,
			},

			args: []string{"path/to/mindnum", "get", "19900719"},
			wantStdOut: `Your mind number is 9!

` + mindnumRepo.DescriptionNine,
			wantStdErr: "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			stdout, stderr, err := tt.fields.capturer.CaptureOutput(
				tt.fields.fnc,
			)
			if err != nil {
				t.Errorf("Capturer.CaptureOutput() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if stdout != tt.wantStdOut {
				t.Errorf("main() : stdout = %v, wantStdOut = %v", stdout, tt.wantStdOut)
			}
			if stderr != tt.wantStdErr {
				t.Errorf("main() : stderr = %v, wantStdErr = %v", stderr, tt.wantStdErr)
			}
		})
	}
}
