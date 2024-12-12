package proxy

import (
	"runtime/debug"
)

type Debug interface {
	ReadBuildInfo() (info *debug.BuildInfo, ok bool)
}

type debugProxy struct{}

func NewDebug() Debug {
	return &debugProxy{}
}

func (*debugProxy) ReadBuildInfo() (info *debug.BuildInfo, ok bool) {
	return debug.ReadBuildInfo()
}
