package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scai-gen",
	Short: "A CLI tool for generating/checking SCAI metadata",
}

var (
	outFile     string
	prettyPrint bool
)

func init() {
	rootCmd.AddCommand(rdCmd)
	rootCmd.AddCommand(assertCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(sigstoreCmd)
	rootCmd.AddCommand(rekorCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
