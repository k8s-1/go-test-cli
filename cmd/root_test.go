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
	input := "jayquellin"
	cmd.SetArgs([]string{"serve", "name", in})
	cmd.Execute()

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != input {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi-via-args", string(out))
	}
}
