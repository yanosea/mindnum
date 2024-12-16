package formatter

import (
	"fmt"

	"github.com/yanosea/mindnum/pkg/errors"
)

// Formatter is an interface that formats the output of mindnum cli.
type Formatter interface {
	Format(result interface{}) string
}

// NewFormatter returns a new instance of the Formatter interface.
func NewFormatter(
	format string,
) (Formatter, error) {
	var f Formatter
	switch format {
	case "text":
		f = NewTextFormatter()
	default:
		return nil, errors.New("invalid format")
	}
	return f, nil
}

// AppendErrorToOutput appends an error to the output.
func AppendErrorToOutput(err error, output string) string {
	if err == nil && output == "" {
		return ""
	}

	var result string
	if err != nil {
		if output == "" {
			result = fmt.Sprintf("Error : %s", err)
		} else {
			result = fmt.Sprintf("Error : %s\n%s", err, output)
		}
	} else {
		result = output
	}

	if result != "" {
		result = Red(result)
	}

	return result
}
