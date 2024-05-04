package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Base command to add new projects or tasks to your Kando",
}

func init() {
	addCmd.AddCommand(addProjectCmd)
	addCmd.AddCommand(addTaskCmd)
}
