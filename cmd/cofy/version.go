package main

import (
	"github.com/overbool/cofy"
	"github.com/spf13/cobra"
)

var all bool

func init() {
	versionCMD.Flags().BoolVar(&all, "all", false, "show all version info")
}

var versionCMD = &cobra.Command{
	Use:   "version [flags]",
	Short: "Show version about app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Cofy version: %s-%s\n", cofy.CurrentVersion, cofy.CurrentCommit)
		if all {
			cmd.Printf("App build date: %s\n", cofy.BuildDate)
			cmd.Printf("System version: %s\n", cofy.Platform)
			cmd.Printf("Golang version: %s\n", cofy.GoVersion)
		}
	},
}
