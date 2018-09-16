package commands

import (
	"fmt"

	"github.com/EdgeSmart/EdgeDeploy/check"
	"github.com/spf13/cobra"
)

var checkCMD = &cobra.Command{
	Use:   "check",
	Short: "Check system environment",
	Long:  `Check system environment`,
	Run: func(cmd *cobra.Command, args []string) {
		err := check.Run()
		if err != nil {
			fmt.Println("Check failed")
			return
		}
		fmt.Println("Check success")
		return
	},
}

func init() {
	addCommand(checkCMD)
}
