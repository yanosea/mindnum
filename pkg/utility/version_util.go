package utility

import (
	"github.com/yanosea/mindnum-pkg/proxy"
)

type VersionUtil interface {
	GetVersion(version string) string
}

type versionUtil struct {
	Debug proxy.Debug
}

func NewVersionUtil(
	debug proxy.Debug,
) VersionUtil {
	return &versionUtil{
		Debug: debug,
	}
}

func (v *versionUtil) GetVersion(version string) string {
	// if version is embedded, return it
	if version != "" {
		return version
	}

	if i, ok := v.Debug.ReadBuildInfo(); !ok {
		return "unknown"
	} else {
		return i.Main.Version
	}
}
