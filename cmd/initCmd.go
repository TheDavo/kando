package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kando/kando"
	"os"
	"path/filepath"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Kando document",
	Run:   initKando,
}

func initKando(cmd *cobra.Command, args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	pathToKando := filepath.Join(homeDir, "kando", "kando.json")
	_, err = os.Stat(pathToKando)

	if err != nil {
		fmt.Println("Kando file does not exist:", pathToKando)
		fmt.Println("Creating Kando file...")
		os.Mkdir(filepath.Join(homeDir, "kando"), os.ModeDir)
		file, err := os.Create(pathToKando)
		if err != nil {
			panic(err)
		}
		file.Close()

		k := kando.NewKando("Personal", pathToKando)
		err = k.Save()
		if err != nil {
			panic(err)
		}
		fmt.Println("Initialized a Kando file!")
		return
	}
	fmt.Println("Kando file already exists at:", pathToKando)
}
