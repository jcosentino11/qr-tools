package cli

import (
	"fmt"
	"io"

	"josephcosentino.me/qr-code-tool/internal/generator"
)

func CmdGen(args []string, stdout, stderr io.Writer) error {
	url := args[1]
	path := args[2]

	err := generator.Generate(url, path)
	if err != nil {
		return fmt.Errorf("failed to generate qr code: %w", err)
	}

	return nil
}
