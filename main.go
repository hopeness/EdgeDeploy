package main

import (
	"fmt"
	"os"

	"github.com/EdgeSmart/EdgeDeploy/commands"
)

func main() {
	command := commands.GetCommand()
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
