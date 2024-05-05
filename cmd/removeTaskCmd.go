package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kando/kando"
	"strconv"
)

var removeTaskCmd = &cobra.Command{
	Use:   "task [id] from [project]",
	Short: "Remove a task from a project",
	Args:  cobra.ExactArgs(3),
	Run:   removeTaskCmdFunc,
}

func removeTaskCmdFunc(cmd *cobra.Command, args []string) {
	if args[1] != "from" {
		fmt.Println("Incorrect usage of command")
		fmt.Println("kando remove task [id] FROM [project]")
		return
	}

	k := kando.Open()
	contains := false
	idArg := args[0]
	projArg := args[2]

	for _, proj := range k.Meta.Projects {
		if projArg == proj {
			contains = true
		}
	}

	if !contains {
		fmt.Println("Error: project", projArg, "not present in Kando file")
		k := kando.Open()
		for projNumber, projTitle := range k.Meta.Projects {
			fmt.Println(projNumber, ",", projTitle)
		}
		return
	}

	idAsInt, err := strconv.Atoi(idArg)
	if err != nil {
		panic(err)
	}
	k.Projects[projArg].RemoveTask(idAsInt)
	err = k.Save()
	if err != nil {
		panic(err)
	}

	fmt.Println("Removing task from project", projArg)
}
