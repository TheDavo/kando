package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"kando/kando"
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
	if args[0] != "to" {
		fmt.Println("Incorrect usage of command")
		fmt.Println("kando add task TO PROJECT")
		return
	}

	k := kando.Open()
	contains := false

	for _, proj := range k.Meta.Projects {
		if args[1] == proj {
			contains = true
		}
	}

	if !contains {
		fmt.Println("Error: project", args[1], "not present in Kando file")
		return
	}

	fmt.Println("Adding task to project", args[1])
	if _, err := tea.NewProgram(
		ui.AddTaskInitialModel(args[1]),
	).Run(); err != nil {
		fmt.Println("could not start program", err)
		os.Exit(1)
	}
}
