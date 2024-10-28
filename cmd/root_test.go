package cmd

import (
	"testing"
)


func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd()

	cmd.SetArgs([]string{"serve", "--name", "jayquellin"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

  // test things
}
