package cmd

import (
	"bytes"
	"io"
	"testing"
)


func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")

	cmd.SetOut(b)

	// Correctly setting the args to use the flag format
	cmd.SetArgs([]string{"serve", "--name", "jayquellin"})

	// Execute the command
	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Read the output
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the output matches the expected value
	if string(out) != "jayquellin" {
		t.Fatalf("expected \"%s\" got \"%s\"", "jayquellin", string(out))
	}
}
