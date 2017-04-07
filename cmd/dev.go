package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Source string

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

	devCmd.PersistentFlags().StringVarP(
		&Source,
		"source",
		"s",
		"tct-api",
		"Specify data source: tct-api, tct-db, cache",
	)
}
