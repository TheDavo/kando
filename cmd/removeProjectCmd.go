package cmd

import (
	"kando/kando"
	"strconv"

	"github.com/spf13/cobra"
)

var removeProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Add a new project to your Kando file",
	Args:  cobra.ExactArgs(1),
	Run:   removeProjectsCmdFunc,
}

func removeProjectsCmdFunc(cmd *cobra.Command, args []string) {

	// Make sure the argument is a string/word, not a number
	if _, err := strconv.Atoi(args[0]); err == nil {
		panic("Arguement must be a word entry")
	}

	k := kando.Open()

	k.RemoveProject(args[0])
	err := k.Save()
	if err != nil {
		panic(err)
	}
}
