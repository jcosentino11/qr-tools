//go:build integration

package tests

import (
	"testing"
)

func TestDecode(t *testing.T) {

	t.Skip("Skipping until decoder is implemented")

	expectedQrCode := "https://josephcosentino.me"
	sampleFilename := "josephcosentino.me.png"

	actualQrCode := DecodeSample(t, sampleFilename)

	if expectedQrCode != actualQrCode {
		t.Fatalf("QR Code does not match. expected: %s, actual: %s", expectedQrCode, actualQrCode)
	}
}
