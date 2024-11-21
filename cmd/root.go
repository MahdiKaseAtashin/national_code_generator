/*
Copyright © 2024 Mahdi Kase Atashin MahdiKaseAtashin@Gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ncg",
	Short: "A national code generator",
	Long: `A comprehensive command-line tool for generating national codes.
		This tool leverages the Cobra CLI library for Go, which simplifies the
		creation and management of command-line applications. Use the 'generate'
		command to create a new national code, with an option to generate a rounded
		code using the -r or --round flag. This application is designed to be efficient
		and user-friendly, providing a seamless experience for generating national codes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
