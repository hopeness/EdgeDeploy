package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deployCMD = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy EdgeDeploy",
	Long:  `Deploy EdgeDeploy`,
	Run: func(cmd *cobra.Command, args []string) {
		// deploy.Run()
		fmt.Println("deploy")
	},
}

func init() {
	addCommand(deployCMD)
}
