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
	"github.com/spf13/cobra"

	"github.com/nyulibraries/dlts-enm/tct"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache responses from TCT API",
	Long: `Cache all TCT API responses needed to run a full reload from cache.
Existing cached files are not cleared first, but are overwritten as new responses
are fetched.  :

	./enm cache
	# Sometime in the future, or on a different machine which has a copy of the
	# generated cache:
	./enm reload --source=cache`,
	Run: func(cmd *cobra.Command, args []string) {
		tct.Source = ReloadSource

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
