package formatter

import (
	"fmt"
	"strings"

	mindnumApp "github.com/yanosea/mindnum/v2/app/application/mindnum"
)

// TextFormatter is a struct that formats the output of mindnum cli.
type TextFormatter struct{}

// NewTextFormatter returns a new instance of the TextFormatter struct.
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

// Format formats the output of mindnum cli.
func (f *TextFormatter) Format(result interface{}) string {
	var formatted string
	switch v := result.(type) {
	case *mindnumApp.VersionUsecaseOutputDto:
		formatted = fmt.Sprintf("mindnum version %s", v.Version)
	case *mindnumApp.GetMindnumUsecaseOutputDto:
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("Your mind number is %d!", v.MindNumber))
		if trimmedDescription := strings.TrimSpace(v.Description); trimmedDescription != "" {
			builder.WriteString("\n\n")
			builder.WriteString(trimmedDescription)
		}
		formatted = builder.String()
	}

	return formatted
}
