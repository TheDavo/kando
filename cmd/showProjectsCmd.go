package cmd

import (
	"fmt"
	"kando/kando"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var showProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Show all the projects titles written in the Kando documents",
	Run:   showProjectsCmdFunc,
}

func showProjectsCmdFunc(cmd *cobra.Command, args []string) {
	// Get the Kando file from the default directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	kandoPath := filepath.Join(homeDir, "kando", "kando.json")
	k := kando.FromFilePath(kandoPath)
	for projNumber, projTitle := range k.Meta.Projects {
		fmt.Println(projNumber, ",", projTitle)
	}
}
