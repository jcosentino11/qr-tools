package cli

import (
	"fmt"
	"io"

	"josephcosentino.me/qr-code-tool/internal/reader"
)

func CmdRead(args []string, stdout, stderr io.Writer) error {
	path := args[1]

	url, err := reader.Read(path)
	if err != nil {
		return err
	}

	_, err = stdout.Write([]byte(url + "\n"))
	if err != nil {
		return fmt.Errorf("write to stdout failed %w", err)
	}

	return nil
}
