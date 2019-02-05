package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nyulibraries/dlts-enm/sitegen"
)

// sitePagesCmd represents the browsetopicslists command
var sitePagesCmd = &cobra.Command{
	Use:   "sitepages",
	Short: "Creates ENM website site pages: About, Home",
	Long: `Creates ENM website site pages: About, Home`,
	Run: func(cmd *cobra.Command, args []string) {
		sitegen.GoogleAnalytics = GoogleAnalytics
		sitegen.Source = SitegenSource
		sitegen.GenerateSitePages(Destination)
	},
}

func init() {
	sitegenCmd.AddCommand(sitePagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sitePagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// browseTopicListsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
