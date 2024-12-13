package utility

import (
	"os"

	"github.com/yanosea/mindnum/pkg/proxy"
)

type Capturable interface {
}

type Capturer struct {
	StdBuffer proxy.Buffer
	ErrBuffer proxy.Buffer
}

func NewCapturer(
	sdtBuffer proxy.Buffer,
	errBuffer proxy.Buffer,
) *Capturer {
	return &Capturer{
		StdBuffer: sdtBuffer,
		ErrBuffer: errBuffer,
	}
}

func (c *Capturer) CaptureOutput(fnc func()) (string, string, error) {
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
