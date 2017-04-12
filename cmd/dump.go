// Copyright © 2017 NYU
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

	"github.com/nyulibraries/dlts/enm/db"
	"github.com/nyulibraries/dlts/enm/tct"
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

		if err := db.DB.Ping(); err == nil {
			fmt.Println( "dump: Database " + db.Database + " responding to ping")
		} else {
			panic("dump: " + err.Error())
		}
	},
}

func init() {
	devCmd.AddCommand(dumpCmd)
}
