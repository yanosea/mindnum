package formatter

import (
	"fmt"
	"strings"

	mindnumApp "github.com/yanosea/mindnum/app/application/mindnum"
)

type TextFormatter struct{}

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

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
