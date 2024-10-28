/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a service",
	Long: `  # Serve a service

e.g. 
go-test-cli serve name=example`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

type serveFlags struct {
  name string
}

var serveArgs = serveFlags{}

func init() {
	rootCmd.AddCommand(serveCmd)

  serveCmd.Flags().StringVarP(&serveArgs.name, "name", "n", "", "name of the service")
}
