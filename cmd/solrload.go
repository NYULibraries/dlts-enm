package cmd

import (
	"fmt"
	"github.com/nyulibraries/dlts-enm/solr"
	"github.com/spf13/cobra"
)

// solrLoadCmd represents the add command
var solrLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads the Solr index",
	Long: `Loads the Solr index from views pages and page_topics`,
	Run: func(cmd *cobra.Command, args []string) {
		err := solr.Init(Server, Port, nil)
		if err != nil {
			panic(fmt.Sprintf("ERROR: couldn't initialize solr: %s\n", err.Error()))
		}

		solr.Source = SolrSource

		solr.Load()
	},
}

func init() {
	solrCmd.AddCommand(solrLoadCmd)
}
