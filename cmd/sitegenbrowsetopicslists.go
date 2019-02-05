package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nyulibraries/dlts-enm/sitegen"
)

// browseTopicsListsCmd represents the browsetopicslists command
var browseTopicsListsCmd = &cobra.Command{
	Use:   "browsetopicslists",
	Short: "Creates ENM website browse topics lists pages",
	Long: `Creates ENM website browse topics lists pages`,
	Run: func(cmd *cobra.Command, args []string) {
		sitegen.GoogleAnalytics = GoogleAnalytics
		sitegen.Source = SitegenSource
		sitegen.GenerateBrowseTopicsLists(Destination)
	},
}

func init() {
	sitegenCmd.AddCommand(browseTopicsListsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// browseTopicsListsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// browseTopicListsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
