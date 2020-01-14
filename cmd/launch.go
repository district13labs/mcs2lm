package cmd

import (
	"github.com/district13labs/mcs2lm/src/launcher"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(launchCmd)
}

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launches an instance on AWS",
	Long:  `Launches a pre-configured instance on AWS`,
	Run: func(cmd *cobra.Command, args []string) {
		launcher.LaunchServer()
	},
}
