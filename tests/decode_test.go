//go:build integration

package tests

import (
	"bytes"
	"testing"

	"josephcosentino.me/qr-code-tool/internal/cli"
)

func TestDecode(t *testing.T) {

	t.Skip("Skipping until decoder is implemented")

	qrCode := "https://josephcosentino.me"
	qrCodeFile := "josephcosentino.me.png"

	var stdout, stderr bytes.Buffer
	args := []string{"qr-decode", SampleImagePath(t, qrCodeFile)}

	err := cli.CmdDecode(args, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Decoding failed: %v", err)
	}

	actualQrCode := stdout.String()

	if actualQrCode != qrCode {
		t.Fatalf("QR Code does not match. expected: %s, actual: %s", qrCode, actualQrCode)
	}
}
