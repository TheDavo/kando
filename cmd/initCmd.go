package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kando/kando"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Kando document",
	Run:   initKando,
}

func initKando(cmd *cobra.Command, args []string) {
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
}
