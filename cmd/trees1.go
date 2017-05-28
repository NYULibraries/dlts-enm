// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"github.com/nyulibraries/dlts-enm/dataviz"
)

// trees1Cmd represents the trees1 command
var trees1Cmd = &cobra.Command{
	Use:   "trees1",
	Short: "Create collapsible topic trees - see NYUP-234",
	Long: `Create collapsible topic trees - see NYUP-234

"ENM Data visualization: create expanding topic tree for Topics"
https://jira.nyu.edu/jira/browse/NYUP-234`,
	Run: func(cmd *cobra.Command, args []string) {
		dataviz.OutputDir = OutputDir
		dataviz.Trees1()
	},
}

func init() {
	datavizCmd.AddCommand(trees1Cmd)
}
