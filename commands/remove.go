package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCMD = &cobra.Command{
	Use:   "remove",
	Short: "Remove edge compute environmental Science",
	Long:  `Remove edge compute environmental Science`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run command")
	},
}

func init() {
	addCommand(removeCMD)
}
