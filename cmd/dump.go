// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/nyulibraries/enm/tct"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Do a print dump of all retrieved JSON data",
	Long: `Print raw umarshaled JSON fetched from ENM TCT or cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		if (Cache) {
			fmt.Println("TODO: fetch JSON from cache and dump it")
		} else {
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
		}
	},
}

func init() {
	devCmd.AddCommand(dumpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
