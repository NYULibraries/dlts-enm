package cmd

import (
	"github.com/spf13/cobra"
)

var Port int
var Server string

// solrCmd represents the solr command
var solrCmd = &cobra.Command{
	Use:   "solr",
	Short: "Manage Solr index",
	Long: `Clears and loads Solr index.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(solrCmd)

	solrCmd.PersistentFlags().StringVarP(
		&Server,
		"server",
		"s",
		"localhost",
		"Solr server",
	)

	solrCmd.PersistentFlags().IntVarP(
		&Port,
		"port",
		"p",
		8983,
		"Solr server port",
	)
}
