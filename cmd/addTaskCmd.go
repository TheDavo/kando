package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	// "kando/kando"
	"kando/ui"
	"os"
	// "path/filepath"
	// "strconv"

	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "task to [project]",
	Short: "Add a task to a project",
	Args:  cobra.ExactArgs(2),
	Run:   addTaskCmdFunc,
}

func addTaskCmdFunc(cmd *cobra.Command, args []string) {
	fmt.Println("Adding task to project", args[1])
	if _, err := tea.NewProgram(
		ui.AddTaskInitialModel(),
	).Run(); err != nil {
		fmt.Println("could not start program", err)
		os.Exit(1)
	}
}
