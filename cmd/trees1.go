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
	// Disabled.  See comments below.  Keeping around for now until we are 100%
	// we'll never revive the original prototype MySQL ENM database.
	_ "github.com/nyulibraries/dlts-enm/dataviz"
	"github.com/spf13/cobra"
)

// trees1Cmd represents the trees1 command
var trees1Cmd = &cobra.Command{
	Use:   "trees1",
	Short: "Create collapsible topic trees - see NYUP-234",
	Long: `Create collapsible topic trees - see NYUP-234

"ENM Data visualization: create expanding topic tree for Topics"
https://jira.nyu.edu/jira/browse/NYUP-234`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("`dataviz trees1` is obsolete.  The last (probably) " +
"working version was https://github.com/NYULibraries/dlts-enm/releases/tag/final-commit-before-moving-to-postgres.")
		fmt.Println("For more details, see https://jira.nyu.edu/jira/browse/NYUP-413.")

		// This is what this command used to do:
		// dataviz.OutputDir = OutputDir
		// dataviz.Trees1()
	},
}

func init() {
	datavizCmd.AddCommand(trees1Cmd)
}
