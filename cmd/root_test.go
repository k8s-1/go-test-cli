package cmd

import (
  "testing"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewRootCmd("hi")
	cmd.Execute()
}
