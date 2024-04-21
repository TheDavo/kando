package cmd

import (
	// "fmt"
	// "os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kando",
	Short: "A task management tool with a kan-do attitude",
	Long: `Kando is a task management tool with plenty of flexibility to make
		task management an easy exercise for the burdened over-worked
		employee.`,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
