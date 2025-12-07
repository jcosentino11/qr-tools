package tests

import "testing"

func TestEncodeDecode(t *testing.T) {
	t.Skip("Skip until encode and decode are implemented")

	expectedQrCode := "https://josephcosentino.me"
	qrCodeFile := "/tmp/josephcosentino.me.png"

	_, _ = Encode(t, expectedQrCode, qrCodeFile)

	actualQrCode := Decode(t, qrCodeFile)

	if expectedQrCode != actualQrCode {
		t.Fatalf("QR Code does not match. expected: %s, actual: %s", expectedQrCode, actualQrCode)
	}
}
