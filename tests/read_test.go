//go:build integration

package tests

import (
	"bytes"
	"testing"

	"josephcosentino.me/qr-code-tool/internal/cli"
)

func TestRead(t *testing.T) {

	t.Skip("Skipping until reader is implemented")

	var stdout, stderr bytes.Buffer
	args := []string{"qr-read", SampleImagePath(t, "josephcosentino.me.png")}

	err := cli.CmdRead(args, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	output := stdout.String()

	expected := "https://josephcosentino.me\n"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
