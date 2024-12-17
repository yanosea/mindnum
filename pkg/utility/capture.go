package utility

import (
	"os"

	"github.com/yanosea/mindnum/pkg/proxy"
)

// Capturable is an interface that captures the output of a function.
type Capturable interface {
	CaptureOutput(fnc func()) (string, string, error)
}

// capturer is a struct that implements the Capturable interface.
type capturer struct {
	// StdBuffer is a buffer for standard output.
	StdBuffer proxy.Buffer
	// ErrBuffer is a buffer for error output.
	ErrBuffer proxy.Buffer
}

// NewCapturer returns a new instance of the capturer struct.
func NewCapturer(
	stdBuffer proxy.Buffer,
	errBuffer proxy.Buffer,
) *capturer {
	return &capturer{
		StdBuffer: stdBuffer,
		ErrBuffer: errBuffer,
	}
}

// CaptureOutput captures the output of a function.
func (c *capturer) CaptureOutput(fnc func()) (string, string, error) {
	origStdout := os.Stdout
	origStderr := os.Stderr
	defer func() {
		os.Stdout = origStdout
		os.Stderr = origStderr
	}()

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr

	fnc()

	wOut.Close()
	wErr.Close()

	if _, err := c.StdBuffer.ReadFrom(rOut); err != nil {
		return "", "", err
	}

	if _, err := c.ErrBuffer.ReadFrom(rErr); err != nil {
		return "", "", err
	}

	stdout := c.StdBuffer.String()
	errout := c.ErrBuffer.String()

	c.StdBuffer.Reset()
	c.ErrBuffer.Reset()

	return stdout, errout, nil
}
