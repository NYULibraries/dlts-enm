package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cache bool

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Temporary command for initial development",
	Long: `Just a temporary command to use during initial development for
	running new code and doing sundry experiments`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev command called")
	},
}

func init() {
	RootCmd.AddCommand(devCmd)

	devCmd.PersistentFlags().BoolVarP(
		&Cache,
		"cache",
		"c",
		false,
		"Get data from cache instead of making live calls to ENM TCT",
	)
}
