package cmd

import (
	"bytes"
	"io"
	"testing"
)

func Test_UpdateCommand(t *testing.T) {
	cmd := updateCmd()

	// Buffer to catch stdout
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)

	// Runs command
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Read and assert Buffer
	out, err := io.ReadAll(buf)
	if err != nil {
		t.Fatal(err)
	}

	expected := "update called\n"
	if string(out) != expected {
		t.Fatalf("expected %q, got %q", expected, string(out))
	}
}
