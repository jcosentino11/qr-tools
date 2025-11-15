package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func SampleImagePath(t *testing.T, filename string) string {
	t.Helper()

	path := filepath.Join("..", "samples", filename)

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("sample image not found: %s (error: %v)", filename, err)
	}

	return path
}
