package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nyulibraries/dlts-enm/tct"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache responses from TCT API",
	Long: `Cache all TCT API responses.  Existing cached files are not cleared first,
but are overwritten as new responses are fetched.  :

	./enm cache`,
	Run: func(cmd *cobra.Command, args []string) {
		tct.Source = "tct-api"

		tctTopics := tct.GetTopicsAll()

		for _, tctTopic := range tctTopics {
			tctTopicDetail := tct.GetTopicDetail(int(tctTopic.ID))

			_ = tctTopicDetail.Basket.TopicHits
		}

		_ = tct.GetIndexPatternsAll()

		tctEpubs := tct.GetEpubsAll()
		for _, tctEpub := range tctEpubs {
			tctEpubDetail := tct.GetEpubDetail(int(tctEpub.ID))

			tctLocations := tctEpubDetail.Locations
			for _, tctLocation := range tctLocations {
				_ = tct.GetLocation(int(tctLocation.ID))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(cacheCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cacheCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cacheCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
