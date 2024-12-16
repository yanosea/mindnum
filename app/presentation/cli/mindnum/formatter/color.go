package formatter

import (
	"github.com/fatih/color"
)

var (
	// Green is a proxy of fatih/color.Green.
	Green = color.New(color.FgGreen).SprintFunc()
	// Red is a proxy of fatih/color.Red.
	Red = color.New(color.FgRed).SprintFunc()
)
