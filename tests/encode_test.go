//go:build integration

package tests

import (
	"testing"
)

func TestEncode(t *testing.T) {
	t.Skip("Skip until encode is implemented")

	expectedQrCode := "https://josephcosentino.me"
	qrCodeFile := "/tmp/qr-code.png"

	_, _ = Encode(t, expectedQrCode, qrCodeFile)

	actualQrCode := DecodeWithExternalLib(t, qrCodeFile)

	if expectedQrCode != actualQrCode {
		t.Fatalf("QR Code does not match. expected: %s, actual: %s", expectedQrCode, actualQrCode)
	}
}
