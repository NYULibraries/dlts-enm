package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var Destination string
var GoogleAnalytics bool
var SitegenSource string

// sitegenCmd represents the sitegen command
var sitegenCmd = &cobra.Command{
	Use:   "sitegen",
	Short: "Generate ENM website components",
	Long: `Generate the following ENM website components:

	* Site pages: About, Home
	* Topics browse lists
	* Topic pages
	`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(sitegenCmd)

	defaultDestination, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	sitegenCmd.PersistentFlags().StringVarP(
		&Destination,
		"destination",
		"d",
		defaultDestination,
		"Filesystem path to write files to",
	)

	sitegenCmd.PersistentFlags().StringVarP(
		&SitegenSource,
		"source",
		"s",
		"database",
		"Specify data source: database, cache",
	)

	sitegenCmd.PersistentFlags().BoolVarP(
		&GoogleAnalytics,
		"google-analytics",
		"g",
		false,
		"Generate Google Analytics code",
	)
}
