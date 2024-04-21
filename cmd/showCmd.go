package cmd

import (
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Base command to show projects or tasks of a project",
}

func init() {
	showCmd.AddCommand(showProjectsCmd)
}
