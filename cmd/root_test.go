package cmd

import (
	"testing"
  "bytes"
  "io"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")

	cmd.SetOut(b)
	cmd.SetArgs([]string{"serve", "name", "jayquellin"})
	cmd.Execute()

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "hi-via-args" {
		t.Fatalf("expected \"%s\" got \"%s\"", "hi-via-args", string(out))
	}
}
