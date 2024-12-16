package presenter

import (
	"os"
	"testing"

	"github.com/yanosea/mindnum/pkg/proxy"
	"github.com/yanosea/mindnum/pkg/utility"
)

func TestPresent(t *testing.T) {
	capturer := utility.NewCapturer(
		proxy.NewBuffer(),
		proxy.NewBuffer(),
	)

	type fields struct {
		fnc      func()
		capturer utility.Capturable
	}
	tests := []struct {
		name       string
		fields     fields
		wantStdOut string
		wantStdErr string
		wantErr    bool
	}{
		{
			name: "positive testing (stdout, output is not empty)",
			fields: fields{
				fnc: func() {
					Present(os.Stdout, "stdout")
				},
				capturer: capturer,
			},
			wantStdOut: "stdout\n",
			wantStdErr: "",
			wantErr:    false,
		},
		{
			name: "positive testing (stdout, output is empty)",
			fields: fields{
				fnc: func() {
					Present(os.Stdout, "")
				},
				capturer: capturer,
			},
			wantStdOut: "",
			wantStdErr: "",
			wantErr:    false,
		},
		{
			name: "positive testing (stderr, output is not empty)",
			fields: fields{
				fnc: func() {
					Present(os.Stderr, "stderr")
				},
				capturer: capturer,
			},
			wantStdOut: "",
			wantStdErr: "stderr\n",
			wantErr:    false,
		},
		{
			name: "positive testing (stderr, output is empty)",
			fields: fields{
				fnc: func() {
					Present(os.Stderr, "")
				},
				capturer: capturer,
			},
			wantStdOut: "",
			wantStdErr: "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, err := tt.fields.capturer.CaptureOutput(
				tt.fields.fnc,
			)
			if err != nil {
				t.Errorf("Capturer.CaptureOutput() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if stdout != tt.wantStdOut {
				t.Errorf("Present() : stdout = %v, wantStdOut = %v", stdout, tt.wantStdOut)
			}
			if stderr != tt.wantStdErr {
				t.Errorf("Present() : stderr = %v, wantStdErr = %v", stderr, tt.wantStdErr)
			}
		})
	}
}
