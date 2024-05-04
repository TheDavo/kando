package cmd

import (
	"fmt"
	// "os"

	"github.com/spf13/cobra"
	"kando/kando"
)

var rootCmd = &cobra.Command{
	Use:   "kando",
	Short: "A task management tool with a kan-do attitude",
	Long: `Kando is a task management tool with plenty of flexibility to make
		task management an easy exercise for the burdened over-worked
		employee.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		kandoPath, existed, err := kando.KandoFileExists(true)

		// error making the Kando file
		if err != nil {
			panic(err)
		}

		if !existed {
			k := kando.NewKando("Personal", kandoPath)
			err = k.Save()
			if err != nil {
				panic(err)
			}
			fmt.Println("Initialized a Kando file!")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(addCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
