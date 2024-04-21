package cmd

import (
	"kando/kando"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var addProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Add a new project to your Kando file",
	Args:  cobra.ExactArgs(1),
	Run:   addProjectsCmdFunc,
}

func addProjectsCmdFunc(cmd *cobra.Command, args []string) {

	// Make sure the argument is a string/word, not a number
	if _, err := strconv.Atoi(args[0]); err == nil {
		panic("Arguement must be a word entry")
	}

	// Get the Kando file from the default directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	kandoPath := filepath.Join(homeDir, "kando", "kando.json")

	k := kando.FromFilePath(kandoPath)

	k.AddProject(args[0])
	err = k.Save()
	if err != nil {
		panic(err)
	}
}
