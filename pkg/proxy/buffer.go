package proxy

import (
	"bytes"
	"io"
)

type Buffer interface {
	ReadFrom(r io.Reader) (int64, error)
	Reset()
	String() string
}

type bufferProxy struct {
	bytes.Buffer
}

func NewBuffer() Buffer {
	return &bufferProxy{}
}

func (b *bufferProxy) ReadFrom(r io.Reader) (int64, error) {
	return b.Buffer.ReadFrom(r)
}

func (b *bufferProxy) Reset() {
	b.Buffer.Reset()
}

func (b *bufferProxy) String() string {
	return b.Buffer.String()
}
