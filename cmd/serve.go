/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve a service",
	Long: `  # Serve a service

e.g. 
go-test-cli serve name=example`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := "Service name = " + serveArgs.name
		return errors.New(err)
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
