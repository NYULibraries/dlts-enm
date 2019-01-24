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

	"github.com/nyulibraries/dlts-enm/sitegen"
)

// topicpagesCmd represents the topicpages command
var topicpagesCmd = &cobra.Command{
	Use:   "topicpages",
	Short: "Creates ENM website topic pages",
	Long: `Creates an ENM website topic page for every topic`,
	Run: func(cmd *cobra.Command, args []string) {
		sitegen.GoogleAnalytics = GoogleAnalytics
		sitegen.Source = SitegenSource
		sitegen.TopicIDs = args
		sitegen.GenerateTopicPages(Destination)
	},
}

func init() {
	sitegenCmd.AddCommand(topicpagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topicpagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topicpagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
