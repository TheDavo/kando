package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Base command to remove projects or tasks from your Kando",
}

func init() {
	removeCmd.AddCommand(removeTaskCmd)
	removeCmd.AddCommand(removeProjectCmd)
}
