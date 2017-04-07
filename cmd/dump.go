package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/nyulibraries/enm/tct"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Do a print dump of all retrieved JSON data",
	Long: `Print raw umarshaled JSON fetched from ENM TCT or cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		tct.Source = Source

		topicsList := tct.GetTopicsAll()
		for i, v := range topicsList {
			fmt.Printf("Topic #%d: %v\n", i, v)
		}

		topicDetail := tct.GetTopicDetail(7399)
		fmt.Println(topicDetail)

		namesList := tct.GetNamesAll()
		for i, v := range namesList {
			fmt.Printf("Name #%d: %v\n", i, v)
		}

		epubsList := tct.GetEpubsAll()
		for i, v := range epubsList {
			fmt.Printf("epub #%d: %v\n", i, v)
		}

		epubDetail := tct.GetEpubDetail(12)
		fmt.Println(epubDetail)

		location := tct.GetLocation(2410)
		fmt.Println(location)

		indexPatternsList := tct.GetIndexPatternsAll()
		for i, v := range indexPatternsList {
			fmt.Print("Index pattern #%d: %v\n", i, v)
		}
	},
}

func init() {
	devCmd.AddCommand(dumpCmd)
}
