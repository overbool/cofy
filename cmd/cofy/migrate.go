package main

import (
	"github.com/spf13/cobra"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
