package commands

import (
	"fmt"

	"github.com/EdgeSmart/EdgeDeploy/install"
	"github.com/spf13/cobra"
)

var installCMD = &cobra.Command{
	Use:   "install",
	Short: "Install docker enging",
	Long:  `Install docker enging`,
	Run: func(cmd *cobra.Command, args []string) {
		err := install.Run()
		if err != nil {
			fmt.Println("Install failed")
			return
		}
		fmt.Println("Install success")
		return
	},
}

func init() {
	addCommand(installCMD)
}
