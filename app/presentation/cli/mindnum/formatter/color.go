package formatter

import (
	"github.com/fatih/color"
)

var (
	// Red is a proxy of fatih/color.Red.
	Red = color.New(color.FgRed).SprintFunc()
)
