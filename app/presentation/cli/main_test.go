package main

import (
	"os"
	"testing"

	"github.com/yanosea/mindnum-pkg/proxy"
	"github.com/yanosea/mindnum-pkg/utility"
)

func Test_main(t *testing.T) {
	capturer := utility.NewCapturer(
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
		capturer *utility.Capturer
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

You are a BALANCER.

BALANCER is like...
The key word in your life is "2".
You are a person who is destined to have two of everything.
You may travel back and forth between the country and abroad, have two homes, or be attracted to two people at the same time.
You have an inquisitive mind and take pleasure in expanding your knowledge in areas you are not familiar with.
They have a keen intuition and sense of perception,
so they may feel pain at events they don't feel comfortable with or at relationships that are only superficial.
Creative workplaces that allow you to express your own opinions are
more suitable for you than simple work or general office work that does not allow you to feel change.
You may get double blessings such as pregnancy and childbirth at the same time of marriage.
`,
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
