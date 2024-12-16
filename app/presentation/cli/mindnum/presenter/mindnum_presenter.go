package presenter

import (
	"fmt"
	"io"
)

// Present is a function that writes the output to the writer.
func Present(writer io.Writer, output string) {
	if output != "" {
		fmt.Fprintf(writer, "%s\n", output)
	}
}
