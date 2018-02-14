// Copyright © 2018 NYU
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

	"github.com/nyulibraries/dlts-enm/sitegen"
)

// browseTopicListsCmd represents the topicpages command
var browseTopicListsCmd = &cobra.Command{
	Use:   "browsetopiclists",
	Short: "Creates ENM website browse topic lists pages",
	Long: `Creates ENM website browse topic lists pages`,
	Run: func(cmd *cobra.Command, args []string) {
		sitegen.Source = SitegenSource
		sitegen.GenerateBrowseTopicLists(Destination)
	},
}

func init() {
	sitegenCmd.AddCommand(browseTopicListsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// browseTopicListsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// browseTopicListsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
